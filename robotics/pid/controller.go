package pid
import (
	"time"
	"github.com/fuzzycow/ev32go/ev3api/device"
	"runtime"
	"log"
	"github.com/fuzzycow/ev32go/robotics/telemetry"
)

const (
	PID_POS_SCALE = 1000
	PID_SPEED_SCALE = 1000
	PID_GAIN_SCALE = 1000
	PID_TIME_SCALE = 1000
)


type Controller struct {
	MoveCh   chan *Move
	interval time.Duration
	motors   []*regMotor
	curMove  *Move
	nextMove *Move
	tl      telemetry.Client
}

func NewController(motors []*device.Motor, k *PID, tl telemetry.Client) *Controller {
	regMotors := make([]*regMotor, len(motors))
	for i, m := range motors {
		regMotors[i] = newRegMotor(m, k)
	}
	return &Controller{
		motors: regMotors,
		interval: 50 * time.Millisecond,
		MoveCh: make(chan *Move),
		tl: tl,
	}
}

type PID struct {
	P, I, D, F int64
}

func (s *PID) reset() {
	s.P = 0
	s.I = 0
	s.D = 0
	s.F = 0
}

type regMotor struct {
	*device.Motor
	max_speed        int
	prev_reg_speed int64
	prev_power	int
	prev_pos int64
	start_pos int64
	start_time time.Time
	prev_pos_error int64
	prev_reg_t       time.Time
	state *PID
	K *PID
	er *PID
}

func newRegMotor(m *device.Motor, k *PID) *regMotor {
	rMotor := &regMotor{
		Motor: m,
		prev_reg_speed: 0,
		max_speed: 900,
		K: k,
		state: &PID{},
	}
	m.SetDutyCycleSP(0)
	m.RunDirect()
	return rMotor
}


func (c *Controller) Start() {
	go c.run()
}

func (c *Controller) Stop() {
	close(c.MoveCh)
}

func (c *Controller) resetReg() {
	for _, m := range c.motors {
		m.start_pos = int64(m.Position())
		m.start_time = time.Now()
		//m.state.reset()
	}
}

func (c *Controller) stopMotors() {
	for _, m := range c.motors {
		m.state.reset()
	}
	zeros := make([]int, len(c.motors))
	c.regulateMotors(zeros)
}

func (c *Controller) getMotorSpeed() []int {
	speed := make([]int, len(c.motors))
	for i,m := range c.motors {
		speed[i] = m.Speed()
	}
	return speed
}

func (c *Controller) run() {
	tick := time.NewTicker(c.interval)

	var moveStartSpeed []int
	speed := make([]int,len(c.motors))

	// reset start time values to something, so that initial dt will be in a sane range
	// c.resetReg()
	var move *Move
	var moveStartTime time.Time

	REG: for {
		select {
		case  <-tick.C:
		case newMove, ok := <-c.MoveCh:
			if !ok {
				break REG
			}
			log.Printf("PID Controller got move: %+v", newMove)
			move = newMove
			moveStartTime = time.Now()
			moveStartSpeed = c.getMotorSpeed()
		}
		if move == nil {
			continue REG
		}
		moving := time.Since(moveStartTime)

		if move.Duration.Nanoseconds() != 0 && moving > move.Duration  {
			if ! move.Fusion {
				c.stopMotors()
			}
			move.Finish()
			move = nil
			continue
		}
		for i,_ := range c.motors {
			speed[i] = move.Speed[i]
			if move.RampUp.Nanoseconds() > 0  {
				if moving < move.RampUp {
					dv:= move.Speed[i]-moveStartSpeed[i]
					left := ( move.RampUp.Seconds() - moving.Seconds() ) / move.RampUp.Seconds()
					//log.Printf("Ramp left: %v",left)
					speed[i] -= int( float64(dv) * left )
				}
			}
		}
		c.regulateMotors(speed)
	}
	c.stopMotors()
}



func (c *Controller) regulateMotors(speed []int) {
	power := make([]int, len(c.motors))
	for i, m := range c.motors {
		p := c.calcPower(m, speed[i])
		power[i] = p
	}
	runtime.LockOSThread()
	for i, m := range c.motors {
		if m.prev_power != power[i] {
			m.SetDutyCycleSP(power[i])
			m.prev_power = power[i]
		}
	}
	runtime.UnlockOSThread()
}


// ./pid_ex -kp 0.1   -ki 0.13
// ./pid_ex -kp 0.1   -ki 0.13 -kd 0.0015


// with 500dt:
// Orig ev3
// ls && ./pid_ex -kp 0.1 -ki 0.006
// ZN
// ls && ./pid_ex -kp 0.133 -ki 0.0036 -kd 0.0009
// No overshoot
// ./cxffgffpid_ex -kp 0.044 -ki 0.0036 -kd 0.0024

func durationMs(t time.Duration) int64 {
	return t.Nanoseconds()/1000000
}

// ./pid_ex -kp 100  -ki 130 -kff 120
func (c *Controller) calcPower(m *regMotor, speed int) int {
	var regSpeed int64
	now := time.Now()
	dtMs := durationMs( now.Sub(m.prev_reg_t))
	intervalMs := durationMs(c.interval)
	if dtMs <= 0 || dtMs > intervalMs * 10 {
		dtMs = intervalMs
	}

	curSpeed := int64(m.Speed())
	position := int64(m.Position()) * PID_POS_SCALE

	regSpeed = int64(normInt(speed, m.max_speed))

	if regSpeed != m.prev_reg_speed {
		m.start_pos = position
		m.start_time = now
	}

	tlm := c.tl.NewMessage("robot," + m.Id())

	runningMs := now.Sub(m.start_time).Nanoseconds() / 1000000

	// FIXME for speed changes?
	speed_error := regSpeed - curSpeed

	pos_error := m.prev_pos + m.prev_reg_speed*dtMs - position
	//pos_error := m.prev_reg_speed*dtMs/1000 - (position - m.prev_pos)
	// pos_error := m.prev_pos + regSpeed*dtMs - position + m.prev_pos_error
	//pos_error_avg := pos_error_from_start * dtMs / runningMs
	pos_error_from_start := ( m.start_pos  +  m.prev_reg_speed*runningMs ) - position

	tlm.AddInt64("error", speed_error).
		AddInt64("error2", pos_error/PID_POS_SCALE).
		AddInt64("error3", pos_error_from_start/PID_POS_SCALE).
		AddInt64("pos", position/PID_POS_SCALE).
		AddInt64("speed", curSpeed).
		AddInt64("reg_speed",regSpeed).
		AddInt64("dt",dtMs)

	prevI := m.state.I

    	m.state.P =  speed_error // m.prev_pos_error * 3/4 + pos_error * 1/4
	m.state.I = pos_error_from_start // m.state.I + pos_error // pos_error//maxAbsInt64(( m.state.I + pos_error ), pos_error_from_start)
	m.state.D = ( pos_error - m.prev_pos_error )
	m.state.F = regSpeed

	term := &PID{
		P: m.K.P * m.state.P,
		I: m.K.I * m.state.I,
		D: m.K.D * m.state.D,
		F: m.K.F * m.state.F,
	}

	power := int(( term.P + term.I/PID_POS_SCALE + term.D/PID_POS_SCALE + term.F)/ PID_GAIN_SCALE)

	m.prev_pos = position
	m.prev_reg_speed = regSpeed
	m.prev_pos_error = pos_error
	m.prev_reg_t = now

	if ( absInt(power) > 100 ) {
		log.Printf("anti-windup applied: %v",m.state.I)
		m.state.I = prevI
	}

	tlm.AddInt64("TermP", term.P/PID_GAIN_SCALE).
		AddInt64("TermI", term.I/PID_GAIN_SCALE/PID_POS_SCALE).
		AddInt64("TermD", term.D/PID_GAIN_SCALE/PID_POS_SCALE).
		AddInt64("TermF", term.F/PID_GAIN_SCALE).
		AddInt("power",power)

	power = normInt(power, 100)

	// log.Printf("Sending tlm: %s",tlm)
	tlm.Send()

	if regSpeed == 0 {
		return 0
	}
	return power
}

/*
func (c *Controller) calcPower_By_Speed(m *regMotor, regSpeed int) int {
	defer c.tlm.Flush()

	curSpeed := m.Speed()
	position := m.Position()
	c.tlm.PublishInt("speed", curSpeed)


	regSpeed = normInt(regSpeed, m.max_speed)

	speed_error := float64(regSpeed - curSpeed)
	c.tlm.PublishInt("error", int(speed_error))
	c.tlm.PublishInt("pos", position)

	now := time.Now()
	dts := now.Sub(m.prev_reg_t).Seconds()
	dtr := dts / c.interval.Seconds()
	dt := dts*500

	if dt <= 0 || dt > 10 {
		dt = 1
	}
	if dtr <= 0 || dtr > 3 {
		dtr = 1
	}
	s := fmt.Sprintf("%f",dtr)
	c.tlm.Publish("dt",s)

	m.state.P = speed_error // speed_error
	m.state.I = m.state.I + speed_error  * dtr
	m.state.D = (speed_error - m.prev_speed_error)  / dtr
	// dPos := float64(m.prev_pos - position) / dtr

	m.prev_pos = position
	m.prev_speed_error = speed_error
	m.prev_reg_speed = curSpeed
	m.prev_reg_t = now

	term := &PID{
		P: m.Gain.P * m.state.P,
		I: m.Gain.I * m.state.I,
		D: m.Gain.D * m.state.D,
	}
	power := int(term.P + term.I + term.D)

	c.tlm.PublishInt("Kp", int(term.P))
	c.tlm.PublishInt("Ki", int(term.I))
	c.tlm.PublishInt("Kd", int(term.D))

	// log.Printf("wspeed=%d, err=%f, PID=%v, K=%v, Term=%v, wpower %v", regSpeed, speed_error, m.pid, m.K, term, power)

	if ( 100 < absInt(power)) {
		log.Printf("anti-windup applied: %v",m.state.I)
		m.state.I -= speed_error * dtr
	}
	c.tlm.PublishInt("power", power)
	power = normInt(power, 100)

	if regSpeed == 0 {
		return 0
	}
	return power
}
*/