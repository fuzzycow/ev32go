// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
	"github.com/fuzzycow/ev32go/robotics/chassis"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/clip"
	"github.com/fuzzycow/ev32go/ev3api/device"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))

func prepareMotor(port string) *device.Motor {
	m := clip.NewMotor(port)
	//ml.Device.SetTracing(true)
	if err := m.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
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

func prepareChassis() *chassis.WheeledChassis {
	ml := prepareMotor(ev3api.OUTPUT_A)
	defer ml.Close()

	mr := prepareMotor(ev3api.OUTPUT_D)
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
		[]chassis.Wheel{lw, rw},
		chassis.TYPE_DIFFERENTIAL)

	c.SetSpeed(30, 45)
	c.SetAccel(0, 0)
	return c
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func main() {
	log.Printf("Starting")
	flag.Parse()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
