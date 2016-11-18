package main
import (
	"time"
	"github.com/fuzzycow/ev32go/clip"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/robotics/chassis"
)

type ChassisActor struct {
	*BasicActor
	Chassis *chassis.WheeledChassis
}

func main() {
	ir := clip.NewInfraredSensor(ev3api.INPUT_1)

	chassisActor := &ChassisActor{
		NewBasic(),
		chassis.NewHolonomicWheel([]chassis.Wheel{},0),
	}

	blockForward := NewTimedSuppressor()
	defer close(blockForward.Inbox())

	lookAhead := NewPoller(10 * time.Millisecond)
	defer close(lookAhead.Inbox())

	lookAhead.Inbox() <- func(Actor) {
		val := ir.Value()

		if val < 20 {
			blockForward <- func(Actor) {
				blockForward.Suppress(time.Second)
			}
			chassisActor.Inbox() <- func(Actor) {
				chassisActor.Chassis.Travel(-30)
			}
		}
	}
}

