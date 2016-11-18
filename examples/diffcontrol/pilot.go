package main

import (
	"github.com/fuzzycow/ev32go/robotics/chassis"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/clip"
	"log"
	"fmt"
	"time"
	"github.com/fuzzycow/ev32go/ev3api/device"
)

func getMotor(port string) *device.Motor {
	m := clip.NewMotor(port)
	//ml.Device.SetTracing(true)
	if err := m.Open(); err != nil {
		log.Fatalf("Failed to open device: %v",err)
	}
	defer m.Close()
	m.Reset()
	//ml.SetStopCommand(ml.StopCommand_Coast())
	m.SetStopCommand(m.StopCommand_Brake())
	m.SetSpeedRegulationEnabled(m.SpeedRegulation_On())
	m.SetPolarity(m.Polarity_Inversed())
	m.SetRampUpSP(100)
	m.SetRampDownSP(100)
	return m
}

func main() {
	ml := getMotor(ev3api.OUTPUT_A)
	defer ml.Close()

	mr := getMotor(ev3api.OUTPUT_D)
	defer mr.Close()

	lw := &chassis.CarWheel{
		Motor: ml,
		Diameter: 9.2,
		Offset: -8.4,
		GearRatio: 1.6667,

	}
	rw := &chassis.CarWheel{
		Motor: mr,
		Diameter: 9.2,
		Offset: 8.4,
		GearRatio: 1.6667,
	}
	c := chassis.New(
		[]chassis.Wheel{lw,rw},
		chassis.TYPE_DIFFERENTIAL)

	c.SetSpeed(30,45)
	c.SetAccel(0,0)

	//time.Sleep(time.Second * 15)
	log.Println("Travel Fwd")
	c.Travel(30)
	c.WaitComplete()

	//time.Sleep(time.Second * 15)
	log.Println("Travel Back")
	c.Travel(-30)
	c.WaitComplete()

	fmt.Println("Rotate")
	c.Rotate(360*2)
	c.WaitComplete()

	fmt.Println("Arc 360 x - 90")
	c.Arc(30,360)
	c.WaitComplete()

	fmt.Println("Arc - 360 x - 90")
	c.Arc(-30,360)
	c.WaitComplete()

	fmt.Println("Arc 360 x - 90")
	c.Arc(15,-360)
	c.WaitComplete()

	fmt.Println("Arc - 360 x - 90")
	c.Arc(-15,-360)
	c.WaitComplete()

	c.StopMotors()
}


