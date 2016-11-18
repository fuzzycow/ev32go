package main

import (
	"encoding/json"
	"log"
)

type Message struct {
	Cmd 	*Cmd			`json:"cmd,omitempty"`
	Reading map[string]int	`json:"reading,omitempty"`
}

type Travel struct {
	Distance float64 `json:"distance"`
}

type Cmd struct {
	Travel *Travel  `json:"travel,omitempty"`
	Stop *Stop  `json:"travel,omitempty"`
	Rotate *Rotate  `json:"rotate,omitempty"`
	Arc *Arc  `json:"arc,omitempty"`
}

type Stop struct {
}

type Rotate struct {
	Angle	 float64 `json:"angle"`
}

type Arc struct {
	Radius   float64 `json:"radius"`
	Angle    float64 `json:"angle"`
}

func main() {
	m := &Message{
		Cmd: &Cmd {
			Arc: &Arc{
				Angle: 3,
				Radius: 55,
			},
		},
	}
	b,err := json.Marshal(m)
	log.Printf("encoded '%s', error: %v",string(b),err)
}

