package chassis

type Chassis interface {
	SetVelocity(float64,float64,float64)
	SetSpeed(float64,float64)
	SetAccel(float64,float64)
	Travel(float64)
	Rotate(float64)
	Arc(float64,float64)
	Stop()
	StopMotors()
	WaitComplete()
	Close()
}

