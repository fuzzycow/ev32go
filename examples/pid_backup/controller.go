package pid

import (
	"time"
	"github.com/paulmach/go.geo/clustering/helpers"
	"github.com/fuzzycow/ev32go/ev3api/device"
)



/*
import (
	"time"
	"github.com/fuzzycow/ev32go/ev3api/device"
	"runtime"
	"fmt"
	"log"
	"github.com/fuzzycow/ev32go/helpers"
)

type Controller struct {
	MoveCh   chan *Move
	interval time.Duration
	motors   []*regMotor
	curMove  *Move
	nextMove *Move
	tlm      helpers.Publisher
}

func NewController(motors []*device.Motor, k *PID, pub helpers.Publisher) *Controller {
	regMotors := make([]*regMotor, len(motors))
	for i, m := range motors {
		regMotors[i] = newRegMotor(m, k)
	}
	return &Controller{
		motors: regMotors,
		interval: 20 * time.Millisecond,
		MoveCh: make(chan *Move),
		tlm: pub,
	}
}

type PID struct {
	P, I, D float64
}

func (s *PID) reset() {
	s.P = 0
	s.I = 0
	s.D = 0
}

type regMotor struct {
	*device.Motor
	max_speed        int
	prev_speed       int
	prev_speed_error float64
	base_power	float64
	prev_power	int
	prev_reg_t       time.Time
	serr_1 float64
	serr_2 float64
	pid              *PID
	K                *PID
}

func newRegMotor(m *device.Motor, k *PID) *regMotor {
	rMotor := &regMotor{
		Motor: m,
		prev_speed: 0,
		max_speed: 900,
		// pid_k: &pid{P: 150, I:35, D: 10},
		K: k, // &PID{P: 10, I:0, D: 0}
		pid: &PID{},
	}
	return rMotor
}

type Move struct {
	Speed []int
	Time  time.Duration
}

func (move *Move) String() string {
	return fmt.Sprintf("speed=%+v, time=%v", move.Speed, move.Time)
}

func (c *Controller) Start() {
	go c.run()
}

func (c *Controller) Stop() {
	close(c.MoveCh)
}

func (c *Controller) resetReg() {
	for _, m := range c.motors {
		m.pid.reset()
		m.prev_speed_error = 0
		m.serr_1 = 0
		m.serr_2 = 0
		m.base_power = 0
	}
}

func (c *Controller) stopMotors() {
	for _, m := range c.motors {
		m.pid.reset()
	}
	zeros := make([]int, len(c.motors))
	c.regulateMotors(zeros)
}

func (c *Controller) run() {
	tick := time.NewTicker(c.interval)
	numMotors := len(c.motors)
	speed := make([]int, numMotors)

	REG: for {
		select {
		case <-tick.C:
			c.regulateMotors(speed)
		case move, ok := <-c.MoveCh:
			if !ok {
				break REG
			}
			log.Printf("got move: %v", move)
			speed = move.Speed
			c.resetReg()
			c.regulateMotors(speed)
		}
	}
	c.stopMotors()
}

func normInt(want, abslimit int) int {
	switch {
	case want > 0 && want > abslimit:
		return abslimit
	case want < 0 && want < - abslimit:
		return -abslimit
	}
	return want
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

func (c *Controller) calcPower_Lejos(m *regMotor, regSpeed int) int {
	defer c.tlm.Flush()
	c.tlm.PublishInt("want_speed", regSpeed)

	curSpeed := m.Speed()
	speed_error := float64(regSpeed - curSpeed)
	regSpeed = normInt(regSpeed, m.max_speed)

	c.tlm.PublishInt("error", int(speed_error))

	m.serr_1 = m.serr_1 * 1 / 3 + speed_error * 2 / 3; // fast smoothing
	m.serr_2 = m.serr_2 * 3 / 4 + speed_error * 1 / 4; // slow smoothing

	c.tlm.PublishInt("serr_1", int(m.serr_1))
	c.tlm.PublishInt("serr_2", int(m.serr_2))

	now := time.Now()
	dt := now.Sub(m.prev_reg_t).Seconds()
	dtr := dt / c.interval.Seconds()
	c.tlm.PublishInt("dtr",int(dtr))
	if dtr <= 0 || dtr > 10 {
		dtr = 1
	}

	// last good values:  -kp 0.13 -ki 0.06 -kd 0.05
	newPower := m.base_power + m.K.P * m.serr_1 + m.K.D * (m.serr_1 - m.serr_2)/dtr
	m.base_power = m.base_power + m.K.I * (newPower - m.base_power)*dtr

	m.prev_reg_t = now
	c.tlm.PublishInt("power", int(newPower))

	if (m.base_power > 100) {
		log.Printf("anti-windup applied")
		m.base_power = 100
	} else if m.base_power < -100 {
		m.base_power = -100
	}
	if newPower > 100 {
		newPower = 100
	} else if newPower < -100{
		newPower = -100
	}

	if regSpeed == 0 {
		return 0
	}

	return int(newPower)
}

// ./pid_ex -kp 0.1   -ki 0.13
// ./pid_ex -kp 0.1   -ki 0.13 -kd 0.0015


// with 500dt:
// Orig ev3
// ls && ./pid_ex -kp 0.1 -ki 0.006
// ZN
// ls && ./pid_ex -kp 0.133 -ki 0.0036 -kd 0.0009
// No overshoot
// ./pid_ex -kp 0.044 -ki 0.0036 -kd 0.0024

func (c *Controller) calcPower(m *regMotor, regSpeed int) int {
	defer c.tlm.Flush()
	c.tlm.PublishInt("want_speed", regSpeed)

	curSpeed := m.Speed()
	regSpeed = normInt(regSpeed, m.max_speed)

	speed_error := float64(regSpeed - curSpeed)
	c.tlm.PublishInt("error", int(speed_error))
	now := time.Now()
	dts := now.Sub(m.prev_reg_t).Seconds()
	dtr := dts / c.interval.Seconds()
	dt := dts*500
	c.tlm.PublishInt("dtr",int(dtr))

	if dt <= 0 || dt > 10 {
		dt = 1
	}

	m.pid.P = speed_error // speed_error
	m.pid.I = m.pid.I + speed_error * dt
	m.pid.D = (speed_error - m.prev_speed_error) / dt

	m.prev_speed_error = speed_error
	m.prev_speed = curSpeed
	m.prev_reg_t = now

	term := &PID{
		P: m.pid.P * m.K.P,
		I: m.pid.I * m.K.I,
		D: m.pid.D * m.K.D,
	}
	power := int(term.P + term.I + term.D)

	c.tlm.PublishInt("Kp", int(term.P))
	c.tlm.PublishInt("Ki", int(term.I))
	c.tlm.PublishInt("Kd", int(term.D))

	// log.Printf("wspeed=%d, err=%f, PID=%v, K=%v, Term=%v, wpower %v", regSpeed, speed_error, m.pid, m.K, term, power)

	if ( 100 < absInt(power)) {
		log.Printf("anti-windup applied: %v",m.pid.I)
		m.pid.I -= speed_error //* dt
	}
	c.tlm.PublishInt("power", power)
	power = normInt(power, 100)

	if regSpeed == 0 {
		return 0
	}
	return power
}

*/