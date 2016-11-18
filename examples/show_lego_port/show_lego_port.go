package main

import (
	"log"
	"github.com/fuzzycow/ev32go/clip"
	"flag"
)

var portName = flag.String("port","in1","lego port to show")

func main() {
	flag.Parse()
	log.Printf("opening lego port %s",*portName)
	dev := clip.NewLegoPort(*portName)

	if err := dev.Open(); err != nil {
		log.Fatalf("Failed to open lego port %s: %v", *portName,err)
	}
	defer dev.Close()

	log.Printf("Lego Port port name=%v, driver name=%v, mode=%v, modes=%v",
		dev.PortName(),
		dev.DriverName(),
		dev.Mode(),
		dev.Modes())
	if dev.Err() != nil {
		log.Fatalf("device error: %v",dev.Err())
	}
	log.Printf("done")
}
