package chassis


import (
	"github.com/gonum/matrix/mat64"
	"log"
	"math"
	"fmt"
	"runtime"
	"time"
	"github.com/fuzzycow/ev32go/robotics/telemetry/influxtl"
	"github.com/fuzzycow/ev32go/ev3api/device"
	"github.com/fuzzycow/ev32go/robotics/utils"
	"github.com/fuzzycow/ev32go/robotics/pid"
)

const (
	TACHOCOUNT = 0
	MAXSPEED = 1
	ROTATIONSPEED = 2

	TYPE_DIFFERENTIAL = 2
	TYPE_HOLONOMIC = 3
)

type WheeledChassis struct {
	wheels       []Wheel
	dummyWheels  int

	forward      *mat64.Dense
	reverse      *mat64.Dense
	forwardAbs   *mat64.Dense
	reverseAbs   *mat64.Dense

	linearSpeed  float64
	linearAccel  float64
	angularSpeed float64
	angularAccel float64

	ctrl *pid.Controller
}


func New(wheels []Wheel, dim int) *WheeledChassis {
	c := &WheeledChassis{wheels: wheels}
	nWheels := len(wheels)
	if nWheels < dim {
		log.Fatalf("The c must have at least %d motorized wheels", dim)
	}
	if dim == TYPE_DIFFERENTIAL {
		c.dummyWheels = 1
	}
	totalWheels := nWheels + c.dummyWheels
	c.forward = newMatrix(totalWheels, 3)
	//motorSpeed := mat64.NewDense(totalWheels, 1,make([]float64,totalWheels))

	for row := 0; row < nWheels; row++ {
		c.forward.SetRow(row, wheels[row].Factors())
	}
	if (c.dummyWheels == 1) {
		c.forward.SetRow(nWheels, []float64{0, 1, 0});
	}
	logMatrix("Fwd", c.forward)
	c.reverse = newMatrix(totalWheels, 3)
	if err := c.reverse.Inverse(c.forward); err != nil {
		panic(fmt.Errorf("Invalid wheel setup, this robot is not controlable. Check position of the wheels: %v", err))
	}
	c.forwardAbs = matrixAbsCopy(c.forward)
	c.reverseAbs = matrixAbsCopy(c.reverse)

	c.InitPidController()

	return c
}


func (c *WheeledChassis) InitPidController() {
	tl := influxtl.NewClient("udp","192.168.1.10:8089","robot")
	if err := tl.Open(); err != nil {
		log.Fatalf("failed to open udp conn to InfluxDB line protocol socket")
	}

	motors := make([]*device.Motor,c.NumMotors())
	for i := 0; i< c.NumMotors();i++ {
		motors[i] = c.MotorN(i)
	}
	ctrl := pid.NewController(
		motors,
		&pid.PID{P: 100,I: 130,D: 0, F: 130},
		tl)
	c.ctrl = ctrl
	ctrl.Start()
}


func (c *WheeledChassis) NumWheels() int {
	return len(c.wheels)
}

func (c *WheeledChassis) WheelN(n int) Wheel {
	return c.wheels[n]
}

func (c *WheeledChassis) NumMotors() int {
	return c.NumWheels()
}

func (c *WheeledChassis) MotorN(n int) *device.Motor {
	return c.wheels[n].GetMotor()
}

func (c *WheeledChassis) SetPidController(ctrl *pid.Controller) {
	c.ctrl = ctrl
}

func (c *WheeledChassis) SetVelocity(linearSpeed, direction, angularSpeed float64) {
	if c.dummyWheels == 1 && ( int(direction) % 180 != 0 ) {
		log.Fatalln("Invalid direction for differential a robot.")
	}

	tSpeed := matrixMul(
		c.forward,
		cartesianMatrix(
			linearSpeed,
			utils.ToRadians(direction),
			angularSpeed))

	motorAccel := matrixMul(
		c.forwardAbs,
		matrixAbsCopy(
			cartesianMatrix(
				c.linearAccel,
				utils.ToRadians(direction),
				c.angularAccel)))

	deltaSpeed := c.getMotorFnMatrix(func(m *device.Motor) int {return -1 * m.Speed()})
	deltaSpeed.Add(deltaSpeed,tSpeed)
	dt := matrixDivElem(deltaSpeed,motorAccel)
	// FIXME - decelleration / negative values?
	longest := mat64.Max(matrixAbsCopy(dt))

	if longest <= 0.001 {
		log.Printf("FIXME2: no speed change")
		return
	}

	speeds := make([]int,c.NumMotors())

	for i := 0; i < c.NumMotors(); i++ {
		speedFloat := tSpeed.At(i, 0)
		speed := int(speedFloat)
		speeds[i] = speed
		// m := c.MotorN(i)
		// m.SetStopCommand(m.StopCommand_Brake())
		// m.Command_RunDirect()
		log.Printf("SetVelocity: motor[%d] speed=%d, accel=%f", i,speed,longest)
		/* rampTimeMs := int(int(longest*1000))

		m := c.MotorN(i)
		m.SetSpeedSP(speed)
		m.SetRampUpSP(rampTimeMs)
		m.SetRampDownSP(rampTimeMs)
		m.SetStopCommand(m.StopCommand_Coast()) */
	}
	move := pid.NewMove().
		WithSpeed(speeds).
		WithFusion(true).
		WithRampUp( time.Duration(int64(longest*1000)) * time.Millisecond)

	c.ctrl.MoveCh <- move

	// runtime.LockOSThread()
	// defer runtime.UnlockOSThread()
	//c.runForever()

	//logMatrix("SetVelocity motorSpeed", motorSpeed)
}

func (c *WheeledChassis) SetSpeed(linearSpeed, angularSpeed float64) {
	if linearSpeed < 0 || angularSpeed < 0 {
		panic("Speed must be greater then 0")
	}
	c.linearSpeed = linearSpeed
	c.angularSpeed = angularSpeed
}


func (c *WheeledChassis) SetAccel(linearAccel, angularAccel float64) {
	if linearAccel < 0 || angularAccel < 0 {
		panic("Speed must be greater then 0")
	}
	c.linearAccel = linearAccel
	c.angularAccel = angularAccel
}



func (c *WheeledChassis) travelCartesian(xSpeed, ySpeed, angularSpeed float64) {
	c.SetVelocity(math.Sqrt(xSpeed * xSpeed + ySpeed * ySpeed), math.Atan2(ySpeed, xSpeed), angularSpeed);
}

func (c *WheeledChassis) Travel(linearDistance float64) {
	if math.IsInf(linearDistance, 1) || math.IsInf(linearDistance, -1) {
		c.SetVelocity(utils.SigFloat64(linearDistance) * c.linearSpeed, 0, 0)
	} else {
		motorDelta := newMatrix(3, 1)
		motorDelta.Mul(c.forward, toMatrix(linearDistance, 0, 0))

		motorSpeed := newMatrix(3, 1)
		motorSpeed.Mul(c.forward, toMatrix(c.linearSpeed, 0, 0))

		c.setMotors(motorDelta, motorSpeed)
	}
}

func (c *WheeledChassis) Rotate(angularDistance float64) {
	if math.IsInf(angularDistance, 1) || math.IsInf(angularDistance, -1) {
		c.SetVelocity(0, 0, utils.SigFloat64(angularDistance) * c.angularSpeed)
	} else {
		motorDelta := newMatrix(3, 1)
		motorDelta.Mul(c.forward, toMatrix(0, 0, angularDistance))

		motorSpeed := newMatrix(3, 1)
		motorSpeed.Mul(c.forward, toMatrix(0, 0, c.angularSpeed))

		c.setMotors(motorDelta, motorSpeed)
	}
}


func (c *WheeledChassis) Arc(radius, angle float64) {
	if angle == 0 {
		return
	}
	ratio := math.Abs(math.Pi * radius / 180)
	switch {
	case math.IsInf(angle, 1) || math.IsInf(angle, -1):
		if ratio > 1 {
			c.SetVelocity(utils.SigFloat64(angle) * c.linearSpeed,
				0,
				utils.SigFloat64(angle) * c.linearSpeed / ratio)
		} else {
			c.SetVelocity(
				utils.SigFloat64(angle) * c.angularSpeed * ratio,
				0,
				utils.SigFloat64(radius) * c.angularSpeed)
		}
	case math.IsInf(radius, 1) || math.IsInf(radius, -1):
		if angle < 0 {
			c.Travel(math.Inf(1))
		} else {
			c.Travel(math.Inf(-1))
		}
	default:
		displacement := toMatrix(
			utils.SigFloat64(angle) * 2 * math.Pi * math.Abs(radius) * math.Abs(angle) / 360,
			0,
			utils.SigFloat64(radius) * angle)

		var (
			tSpeed    *mat64.Dense
			tAccel *mat64.Dense
		)
		if ratio > 1 {
			tSpeed = toMatrix(c.linearSpeed, 0, c.linearSpeed / ratio)
			tAccel = toMatrix(c.linearAccel, 0, c.linearAccel / ratio)
		} else {
			tSpeed = toMatrix(c.angularSpeed * ratio, 0, c.angularSpeed)
			tAccel = toMatrix(c.angularAccel * ratio, 0, c.angularAccel)
		}

		motorDelta := matrixMul(c.forward, displacement)
		logMatrix("displacement", displacement)
		logMatrix("motor delta", motorDelta)

		// using abs of delta, to fix going backwards
		mRatio := matrixMulEach(
			motorDelta,
			1 / mat64.Max(matrixAbsCopy(motorDelta)))

		logMatrix("mratio", mRatio)
		motorSpeed := matrixMulEach(
			mRatio,
			mat64.Max(matrixMul(c.forwardAbs, tSpeed)))

		//FIXME - Need Abs for Accel ???
		motorAccel := matrixMulEach(
			mRatio,
			mat64.Max(matrixMul(c.forwardAbs, tAccel)))



		_ = motorAccel
		c.setMotors(motorDelta, motorSpeed)
	}
}


func (c *WheeledChassis) Stop() {
	c.SetVelocity(0, 0, 0)
}

func (c *WheeledChassis) StopMotors() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	for i := 0; i < c.NumMotors(); i++ {
		c.MotorN(i).Stop()
	}
}

func (c *WheeledChassis) Close() {
	c.StopMotors()

	for i := 0; i < c.NumMotors(); i++ {
		c.MotorN(i).Close()
	}
}

func (c *WheeledChassis) runForever() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	for i := 0; i < c.NumMotors(); i++ {
		c.MotorN(i).RunForever()
	}
}

func (c *WheeledChassis) runToRelPos() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	for i := 0; i < c.NumMotors(); i++ {
		c.MotorN(i).RunToRelPos()
	}
}

func (c *WheeledChassis) WaitComplete() {
	for i := 0; i < c.NumMotors(); i++ {
		WAIT: for {
			state := c.MotorN(i).GetAttrString("state")
			if state == "" || state == "holding" {
				break WAIT
			} else {
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

func (c *WheeledChassis) setMotors(motorDelta, motorSpeed mat64.Matrix) {
	for i := 0; i < c.NumMotors(); i++ {
		//motor := c.wheels[i].motor
		speed := int(motorSpeed.At(i, 0))
		delta := int(motorDelta.At(i, 0))
		fmt.Printf("Motor %d delta=%d, speed=%d\n", i, delta, speed)
		m := c.MotorN(i)

		m.SetStopCommand(m.StopCommand_Brake())
		m.SetSpeedRegulationEnabled(m.SpeedRegulation_On())
		m.SetPolarity(m.Polarity_Inversed())
		m.SetRampUpSP(500)
		m.SetRampDownSP(500)

		m.SetSpeedSP(speed)
		m.SetPositionSP(delta)
		m.SetStopCommand(m.StopCommand_Brake())
	}
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	c.runToRelPos()
}

type motorAttrGetFn func(*device.Motor) int

// WARNING: Will misfire is called for a non-int attr
func (c *WheeledChassis) getMotorFnMatrix(fn motorAttrGetFn) *mat64.Dense {
	values := make([]float64,c.NumMotors())
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for i := 0; i < c.NumMotors(); i++ {
		values[i] = float64(fn(c.MotorN(i)))
	}
	return mat64.NewDense(3,1,values)
}







