package main

import (
	"github.com/fuzzycow/ev32go/robotics/pid"
	"github.com/fuzzycow/ev32go/ev3api/device"
	"github.com/fuzzycow/ev32go/clip"
	"time"
	"log"
	"github.com/fuzzycow/ev32go/robotics/telemetry/influxtl"
	"flag"
)

var kp = flag.Int64("kp", 0, "Kp")
var ki = flag.Int64("ki", 0, "Ki")
var kd = flag.Int64("kd", 0, "Kd")
var kff = flag.Int64("kff", 0, "Kff")

func main() {
	flag.Parse()
	tl := influxtl.NewClient("udp","192.168.1.10:8089","robot")
	if err := tl.Open(); err != nil {
		log.Fatalf("failed to connect to mqtt")
	}
	defer tl.Close()

	m := clip.NewMotor("outA")
	defer m.Stop()
	m.Open()
	m.Reset()
	m.RunDirect()
	m.SetTracing(false)
	c := pid.NewController(
		[]*device.Motor{m},
		&pid.PID{P: *kp,I: *ki,D: *kd, F: *kff},
		tl)

	c.Start()
	defer c.Stop()


	speeds := []int{100,300,500}
	for _,speed := range speeds {
		log.Printf("starting move with speed %d",speed)
		move := pid.NewMove().
			WithSpeed([]int{speed}).
			WithDuration(7* time.Second).
			WithFusion(true)

		c.MoveCh <- move
		<- move.Done()
	}

	/*
	accel := 100
	for _,speed := range speeds {
		log.Printf("starting move with speed %d,accel %d",speed,accel)
		move := pid.NewMove().
			WithSpeed([]int{speed}).
			WithDuration(5* time.Second).
			WithFusion(true).
			WithRampUp(2*time.Second)

		c.MoveCh <- move
		<- move.Done()
	}*/
}

// ./pid_ex -kp 0.175 -ki 0.1 -kd 0.0375

