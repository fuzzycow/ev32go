package main

import (
	"log"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/clip"
	"time"
	"github.com/fuzzycow/ev32go/ev3api/device"
)

func testManualAccel(m *device.Motor) {
	m.Reset()
	log.Printf("manual accel test")
	m.SetSpeedRegulationEnabled("on")
	m.SetStopCommand(m.StopCommand_Coast())
	for i:=1;i<6;i++ {
		speed := i * 100
		log.Printf("setting speed to %d",speed)
		m.SetSpeedSP(speed)
		m.RunForever()
		time.Sleep(5*time.Second)

	}
	m.Stop()
}

func testAutoAccel(m *device.Motor) {
	m.Reset()
	log.Printf("auto accel test")
	m.SetSpeedRegulationEnabled("on")

	dt := time.Second * 5
	m.SetStopCommand(m.StopCommand_Coast())

	for i:=1;i<6;i++ {
		speed := i * 100
		m.SetRampUpSP(int(dt/time.Millisecond))
		log.Printf("setting speed to %d",speed)
		m.SetSpeedSP(speed)
		m.RunForever()
		time.Sleep(dt)
	}
	m.Stop()
}



func main() {
	log.Printf("started")
	m := clip.NewMotor(ev3api.OUTPUT_A)
	if err := m.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	defer m.Close()
	m.Reset()

	testManualAccel(m)
	testAutoAccel(m)
}
