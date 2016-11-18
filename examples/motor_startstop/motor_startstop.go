package main

import (
	"log"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/clip"
	"time"
	"github.com/fuzzycow/ev32go/helpers/monitor"
)


func main() {
	log.Printf("started")
	m := clip.NewMotor(ev3api.OUTPUT_A)
	if err := m.Open(); err != nil {
		log.Fatalf("Failed to open device: %v",err)
	}
	defer m.Close()
	m.Reset()

	touch := clip.NewSensor(ev3api.INPUT_1)
	if err := touch.Open(); err != nil {
		log.Fatalf("Failed to open device: %v",err)
	}
	defer touch.Close()

	mon := monitor.New(monitor.Changes, 100 * time.Millisecond, 60 * time.Second )
	defer mon.Stop()

	motorCh := mon.PollBool(func() bool { return m.Speed() != 0 } )
	touchCh := mon.PollBool(func() bool { return touch.ValueN(0) == 1 })

	log.Printf("reading monitor...")
	MONITORING: for {
		select {
		case mstate, ok := <- motorCh:
			if !ok {
				break MONITORING
			}
			log.Printf("motor new status: %v", mstate)
		case tstate, ok := <- touchCh:
			if !ok {
				break MONITORING
			}
			if tstate {
				log.Printf("touch sensor pressed - running mottor")
				m.SetDutyCycleSP(50)
				m.SetTimeSP(10000)
				m.RunTimed()
			} else {
				log.Printf("touch sensor released - stopping mottor")
				m.Stop()
			}
		}
	}
}
