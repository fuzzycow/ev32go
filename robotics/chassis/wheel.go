package chassis

import (
	"github.com/fuzzycow/ev32go/ev3api/device"
	"math"
)

type Wheel interface {
	GetMotor() *device.Motor
	Factors() []float64
}

type CarWheel struct {
	Motor *device.Motor
	Diameter float64
	GearRatio float64
	Offset float64
}

func (w *CarWheel) GetMotor() *device.Motor {
	return w.Motor
}

func (w *CarWheel) Factors() []float64 {
	return []float64{
		360 * w.GearRatio / (w.Diameter * math.Pi),
		0,
		- (2.0 * w.Offset * w.GearRatio / w.Diameter),
	}
}

