package chassis

import (
	"github.com/fuzzycow/ev32go/ev3api/device"
	"github.com/fuzzycow/ev32go/robotics/nav"
	"github.com/fuzzycow/ev32go/robotics/utils"
	"math"
)

type HolonomicWheel struct {
	Motor *device.Motor
	Diameter float64
	GearRatio float64
	Pose *nav.Pose
}

func NewHolonomicWheel(motor *device.Motor, diameter float64) *HolonomicWheel {
	w := &HolonomicWheel{
		Pose: nav.NewPose(0,0,0),
		Motor: motor,
		Diameter: diameter,
		GearRatio: 1}
	return w

}

func (w *HolonomicWheel) GetMotor() *device.Motor {
	return w.Motor
}

func (w *HolonomicWheel) Factors() []float64 {
	return []float64{
		math.Cos(utils.ToRadians(w.Pose.Heading)) * 360 / ( w.Diameter * math.Pi * w.GearRatio ),
		math.Sin(utils.ToRadians(w.Pose.Heading)) * 360 / ( w.Diameter * math.Pi * w.GearRatio ),
		2 * w.Pose.Length() / (w.Diameter * w.GearRatio )}
}


func (w *HolonomicWheel) PolarPosition(angle, radius float64) *HolonomicWheel {
	p := nav.NewPose(
		radius * math.Cos(utils.ToRadians(angle)),
		radius * math.Sin(utils.ToRadians(angle)),
		angle)
	p.RotateUpdate(90)

	w.Pose = p
	return w
}

func (w *HolonomicWheel) CartesianPosition(x,y,angle float64) *HolonomicWheel {
	w.Pose = nav.NewPose(x,y,angle)
	return w
}
