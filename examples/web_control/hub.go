// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
import (
	"github.com/fuzzycow/ev32go/robotics/chassis"
	"log"
	"fmt"
)
type Cmd struct {
	Name     string `json:"name"`
	Distance float64 `json:"distance"`
	Radius   float64 `json:"radius"`
	Angle    float64 `json:"angle"`
}


// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	chassis     *chassis.WheeledChassis
	commands    chan *Cmd

	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast   chan []byte

	// Register requests from the connections.
	register    chan *connection

	// Unregister requests from connections.
	unregister  chan *connection

}

var h = hub{
	broadcast:   make(chan []byte),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
	commands:    make(chan *Cmd),
}

func (h *hub) announceCmd(prefix string, cmd *Cmd) {
	msg := prefix + fmt.Sprintf(" %+v", cmd)
	log.Println(msg)
	//h.broadcast <- []byte(msg)
}

func (h *hub) applyCmd(cmd *Cmd) {
	log.Printf("got cmd %+v", cmd)
	switch {
	case cmd.Name == "stop":
		h.chassis.Stop()
		h.announceCmd("cmd", cmd)
	case cmd.Name == "travel":
		h.chassis.Travel(cmd.Distance)
		h.announceCmd("cmd", cmd)
	case cmd.Name == "rotate":
		h.chassis.Rotate(cmd.Angle)
		h.announceCmd("cmd", cmd)
	case cmd.Name == "arc":
		h.chassis.Arc(cmd.Radius, cmd.Angle)
		h.announceCmd("cmd", cmd)
	default:
		h.announceCmd("bad command: ", cmd)
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m, ok := <-h.broadcast:
			if !ok {
				log.Printf("broadcast chan closed")
			}
			log.Printf("broadcast request %s" + string(m))
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					log.Printf("could not send to conn - deleting it")
					close(c.send)
					delete(h.connections, c)
				}
			}
		case cmd := <-h.commands:
			log.Print("got commad - applying")
			h.applyCmd(cmd)
		}
	}
}
