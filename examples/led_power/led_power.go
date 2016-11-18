package main

import (
	"log"
	"github.com/fuzzycow/ev32go/clip"

)


func main() {
	log.Printf("started")
	psu := clip.NewPowerSupply()
	if err := psu.Open(); err != nil {
		log.Fatalf("Failed to open power supply: %v", err)
	}
	defer psu.Close()

	log.Printf("PSU current=%v, voltage=%v",psu.MeasuredCurrent(),psu.MaxVoltage())

	ledName := "ev3-left0:red:ev3dev"
	led := clip.NewLED(ledName)
	if err := led.Open(); err != nil {
		log.Fatalf("Failed to open LED %s: %v",ledName, err)
	}
	defer led.Close()

	log.Printf("LED %s brightness=%v, trigger=%v",led,led.Brightness(),led.Trigger())

	log.Printf("done")
}
