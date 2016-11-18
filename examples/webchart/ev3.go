package main
import (
	"github.com/fuzzycow/ev32go/robotics/chassis"
	"log"
	"fmt"
)

var Chassis *chassis.WheeledChassis


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

