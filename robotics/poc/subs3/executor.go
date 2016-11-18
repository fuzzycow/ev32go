package main

import (
	"log"
	"time"
)

type Executor struct {
	inbox InboxCh
	done DoneCh
	stop DoneCh
	fn func(Actor)
}

func NewExecutor(pri int, fn func()) *Executor {
	return &Executor{
		inbox: make(InboxCh),
		done: make(DoneCh),
		stop: make(DoneCh),
	}
}

func (actor *Executor) Inbox() InboxCh { return actor.inbox }
func (actor *Executor) Done() DoneCh { return actor.done }
func (actor *Executor) Pri() int { return actor.Pri() }
func (actor *Executor) Stop() {
	close(actor.stop)
	<-actor.done
}

func (actor *Executor) Act() {
	ticker := time.NewTicker(time.Second)

	LOOP: for {
		select {
		case guest, ok := <-actor.Inbox():
			if !ok {
				break LOOP
			}
			actor.fn(guest)
		case <- ticker.C:
			log.Printf("tick")
		}
	}
}
