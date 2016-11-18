package main

import (
	"log"
	"github.com/fuzzycow/ev32go/clip"
	"time"
	"flag"
)

var portName = flag.String("port", "in1", "lego port to show")

func main() {
	flag.Parse()
	log.Printf("opening lego port %s",*portName)
	dev := clip.NewSensor(*portName)
	if err := dev.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	defer dev.Close()

	for i:=0;i<10;i++ {
		log.Printf("sensor %s value0: %v", dev.PortName(), dev.ValueN(0))
		time.Sleep(time.Second)
	}
}
