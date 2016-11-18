package main

import (
	"time"
)

type Poller struct {
	*BasicActor
	interval time.Duration
}

func NewPoller(interval time.Duration) *Poller {
	return &Poller{
		interval: interval,
	}
}

func (actor *Poller) Start() {
	var fn func(Actor)
	var ok bool
	ticker := time.NewTicker(actor.interval)
	defer ticker.Stop()

	ACT: go func() {
		select {
		case fn, ok = <- actor.Inbox():
			if !ok {
				break ACT
			}
		case <-ticker:
			fn(actor)
		}
	}()
}


