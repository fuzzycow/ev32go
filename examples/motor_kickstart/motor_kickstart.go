package main

import (
	"log"
	"time"
	"github.com/fuzzycow/ev32go/clip"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/helpers/monitor"
)

func signint(n int) int {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	default:
		return 0
	}
}

func absint(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func main() {
	log.Printf("started")
	m := clip.NewMotor(ev3api.OUTPUT_A)
	if err := m.Open(); err != nil {
		log.Fatalf("Failed to open device: %v",err)
	}
	defer m.Close()
	m.Reset()
	m.SetStopCommand(m.StopCommand_Coast())

	mon := monitor.New(monitor.Changes, 100 * time.Millisecond, 120 * time.Second )
	defer mon.Stop()

	kickstartCh := mon.PollInt(func() int {
		speed := m.Speed()
		dc := m.DutyCycle()
		if dc == 0 && absint(speed) > 30 {
			return speed
		}
		return 0
	} )

	stateCh := mon.PollString(func() string {
		return m.GetAttrString("state")
	})

	var ksc int = 0

	log.Printf("start turning the motor manually, to 'kickstart' it")
	MONITORING: for {
		select {
		case ks, ok := <-kickstartCh:
			if !ok {
				break MONITORING
			}
			if absint(ks) > 0 {
				ksc += 1
			}
			if ksc > 2 {
				ksc = 0
				log.Printf("Kickstarted with speed=%v",ks)
				m.SetDutyCycleSP(50 * signint(ks))
				m.SetTimeSP(5 * ks)
				m.RunTimed()
			}
		case state := <- stateCh:
			log.Printf("Motor state: [%v] ",state)
		}
	}
	m.Stop()
}
