package main

import (
	"github.com/fuzzycow/ev32go/robotics/chassis"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/ev3api/device/sensor"
	"github.com/fuzzycow/ev32go/clip"
	"github.com/fuzzycow/ev32go/ev3api/device"
	"github.com/fuzzycow/ev32go/helpers/monitor"
	"log"
//"fmt"
	"time"
	"math"
)

const (
	IR_NONE = iota
	IR_RED_UP
	IR_RED_DOWN
	IR_BLUE_UP
	IR_BLUE_DOWN
	IR_RED_UP_BLUE_UP
	IR_RED_UP_BLUE_DOWN
	IR_RED_DOWN_BLUE_UP
	IR_RED_DOWN_BLUE_DOWN
	IR_BEACON
	IR_RED_UP_RED_DOWN
	IR_BLUE_UP_BLUE_DOWN
)

var c *chassis.WheeledChassis
var ir *sensor.Infrared


func prepareMotor(port string) *device.Motor {
	m := clip.NewMotor(port)
	//ml.Device.SetTracing(true)
	if err := m.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	m.Reset()
	//ml.SetStopCommand(ml.StopCommand_Coast())
	//m.SetStopCommand(m.StopCommand_Brake())
	m.SetPolarity(m.Polarity_Inversed())
	return m
}

func prepareChassis() *chassis.WheeledChassis {
	mD := prepareMotor(ev3api.OUTPUT_D)
	mA := prepareMotor(ev3api.OUTPUT_A)
	mB := prepareMotor(ev3api.OUTPUT_B)

	w1 := chassis.NewHolonomicWheel(mD, 4.8).PolarPosition(0, 15 * 0.8)
	w2 := chassis.NewHolonomicWheel(mA, 4.8).PolarPosition(120, 15 * 0.8)
	w3 := chassis.NewHolonomicWheel(mB, 4.8).PolarPosition(240, 15 * 0.8)
	c = chassis.New([]chassis.Wheel{w1, w2, w3}, chassis.TYPE_HOLONOMIC)
	return c
}

func prepareIR() *sensor.Infrared {
	ir := clip.NewInfraredSensor(ev3api.INPUT_1)
	if err := ir.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	return ir
}


func handleIRCmd(n int) {
	c.SetSpeed(20, 60)
	switch {
	case n == IR_NONE:
	// do nothing
	case n == IR_RED_UP:
		c.Arc(-30, 90)
	case n == IR_RED_DOWN:
		c.Arc(-30, -90)
	case n == IR_BLUE_UP:
		c.Arc(30, 90)
	case n == IR_BLUE_DOWN:
		c.Arc(30, -90)
	case n == IR_RED_UP_BLUE_UP:
		c.Travel(30)
	case n == IR_RED_UP_BLUE_DOWN:
		c.Rotate(-90)
	case n == IR_RED_DOWN_BLUE_UP:
		c.Rotate(90)
	case n == IR_RED_DOWN_BLUE_DOWN:
		c.Travel(-30)
	case n == IR_BEACON:
		log.Printf("Breacon on ")
	case n == IR_RED_UP_RED_DOWN:
		// c.Rotate(-180)
		c.Stop()
	case n == IR_BLUE_UP_BLUE_DOWN:
		c.Stop()
	default:
		log.Printf("Unexpected IR reading %v", n)
	}
}


func pilot() {
	log.Printf("Started")

	c = prepareChassis()
	defer c.Close()
	ir = prepareIR()
	defer ir.Close()



	log.Printf("devices initialized")

	mon := monitor.New(monitor.Changes, 100 * time.Millisecond, 180 * time.Second)
	defer mon.Stop()

	val0_Ch := mon.PollInt(func() int { return ir.ValueN(0) })
	val1_Ch := mon.PollInt(func() int { return ir.ValueN(1) })

	var (
		val0 int = 0
		val1 int = 0
		ok bool
	)

	ir.SetMode(ir.Mode_IR_SEEK())

	seekInterval := time.Millisecond * 300
	chaseTicker := time.NewTicker(seekInterval)
	mayChase := true

	lastSeen := 0
	lastIrHeading := 0

	log.Printf("reading monitor...")

	MONITORING: for {
		select {
		case val0, ok = <-val0_Ch:
			if !ok {
				break MONITORING
			}
		case val1, ok = <-val1_Ch:
			if !ok {
				break MONITORING
			}
		case <-chaseTicker.C:
			mayChase = true
		}


		switch {
		case val0 == 0 && val1 == -128:
			// log.Printf("Lost beacon !")
			c.SetSpeed(30, 30)
			c.SetAccel(30, 90)
			c.Rotate(math.Inf(lastSeen))
			continue MONITORING
		//case val0 == 0 && val1 == 100:
		case mayChase:
			log.Printf("SEEK MODE: Handling %d / %d",val0,val1)

			irHeading := val0
			irDist := val1

			c.SetSpeed(30, 60)
			c.SetAccel(45, 60)

			var sign int = 1
			if irHeading < 0 && lastIrHeading < 0 && irHeading < lastIrHeading {
				sign = -1
			}
			if irHeading > 0 && lastIrHeading > 0 && irHeading > lastIrHeading {
				sign = -1
			}


			heading := float64(irHeading) * 4.5
			linearSpeed := 30 * float64(irDist) / 100
			angularSpeed := 60 * float64(sign * irHeading) / 25

			log.Printf("Set velocity linSpeed=%f, heading=%f, angSpeed=%f",linearSpeed,heading,angularSpeed)
			c.SetVelocity(linearSpeed,heading,angularSpeed)
			mayChase = false
		default:
			lastIrHeading = val1

		}
	}
	log.Printf("DONE")
}


func main() {
	pilot()
}

