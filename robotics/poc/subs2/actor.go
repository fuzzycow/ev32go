package main

import (
	"time"
	"log"
)

type Actor interface {
	Inbox() InboxCh
	Start()
	Stop()
	String() string
	Done() DoneCh
}

type InboxCh chan func(Actor)
type DoneCh chan struct {}



type Basic struct {
	inbox InboxCh
}

func NewBasic() *Basic {
	return &Basic{
		inbox: make(InboxCh),
	}
}

func (actor *Basic) Inbox() InboxCh { return actor.inbox }
func (actor *Basic) Stop() {
	close(actor.Inbox())
}

func (actor *Basic) Start() {
	go func() {
		for fn := range actor.Inbox() {
			fn(actor)
		}
	}()
}

type TimedSuppressor struct {
	*Basic
	expires time.Time
	timeout chan <-time.Time
}

func NewTimedSuppressor() *TimedSuppressor {
	return &TimedSuppressor{
		expires: time.Time{},
	}
}

func (actor *TimedSuppressor) Suppress(period time.Duration) {
	postpone := time.Now().Add(period)
	if postpone.After(actor.expires) {
		actor.expires = postpone
		actor.timeout = time.After(period)
	}
}

func (actor *TimedSuppressor) Start() {
	ACT: go func() {
		select {
		case fn, ok := actor.Inbox():
			if !ok {
				break ACT
			}
		// BUG
			if actor.timeout == nil {
				fn(actor)
			}
		case <-actor.timeout:
			actor.timeout = nil
		}
	}()
}



type Poller struct {
	*Basic
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




