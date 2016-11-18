package main

import (
	"time"
	"sync"
)

type TimedSuppressor struct {
	*BasicActor

	mu sync.Mutex
	expires time.Time
}

func NewTimedSuppressor() *TimedSuppressor {
	return &TimedSuppressor{
		expires: time.Time{},
	}
}

func (actor *TimedSuppressor) Suppress(period time.Duration) {
	actor.mu.Lock()
	defer actor.mu.Unlock()
	postpone := time.Now().Add(period)
	if postpone.After(actor.expires) {
		actor.expires = postpone
	}
}

func (actor *TimedSuppressor) Start() {
	ACT: go func() {
		select {
		case mesg, ok := <- actor.Inbox():
			if !ok {
				break ACT
			}
			actor.mu.Lock()
			to := actor.expires.Before(time.Now())
			actor.mu.Unlock()
			if to == nil {
				mesg.Pri()
			}
		}
	}()
}





