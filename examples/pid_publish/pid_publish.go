package main

import (

	"github.com/fuzzycow/ev32go/clip"
	"time"
	"log"
	"github.com/fuzzycow/ev32go/robotics/telemetry/influxtl"
	"flag"
)

var kp = flag.Float64("kp", 0, "Kp")
var ki = flag.Float64("ki", 0, "Ki")
var kd = flag.Float64("kd", 0, "Kd")

func main() {
	flag.Parse()
	tl := influxtl.NewClient("udp","192.168.1.10:8089","robot")
	if err := tl.Open(); err != nil {
		log.Fatalf("failed to connect to mqtt")
	}
	defer tl.Close()

	m := clip.NewMotor("outA")
	defer m.Stop()
	m.SetTracing(false)
	m.Open()
	m.Reset()

	tlm := tl.NewMessage("")

	go func() {
		tick := time.NewTicker(time.Millisecond * 20)
		for range tick.C {
			p := m.DutyCycle()
			tlm.AddInt("power",p).Send()
		}
	}()
	time.Sleep(time.Second)
	m.SetSpeedRegulationEnabled(m.SpeedRegulation_On())
	m.SetSpeedSP(300)
	m.SetTimeSP(10*1000)
	m.RunTimed()
	time.Sleep(time.Second*15)
}

// ./pid_ex -kp 0.175 -ki 0.1 -kd 0.0375
