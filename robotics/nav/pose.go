package nav

import (
	"fmt"
	"github.com/paulmach/go.geo"
	"math"
	"github.com/fuzzycow/ev32go/robotics/utils"
)

type Pose struct {
	Loc *geo.Point
	Heading float64
}


func NewPose(x,y,heading float64) *Pose {
	p := &Pose{
		Loc: geo.NewPoint(x,y),
		Heading: heading,
	}
	return p
}

func (p *Pose) RotateUpdate( angle float64) {
	p.Heading += angle
	for p.Heading < 180 {
		p.Heading += 360
	}
	for p.Heading > 180 {
		p.Heading -= 360
	}
}

func (p *Pose) MoveUpdate(distance float64) {
	p2 := geo.NewPoint(
		distance * math.Cos(utils.ToRadians(p.Heading)),
		distance * math.Sin(utils.ToRadians(p.Heading)),
	)
	p.Loc.Add(p2)
}

func (p *Pose) Length() float64 {
	return math.Sqrt(p.Loc.X() * p.Loc.X() + p.Loc.Y() * p.Loc.Y())
	// zero := geo.NewPoint(0,0)
	// return p.Loc.DistanceFrom(zero)
}

func (p *Pose) String() string {
	return fmt.Sprintf("Pose X=%f, Y=%f, H=%f",p.Loc.X(),p.Loc.Y(),p.Heading)

}
