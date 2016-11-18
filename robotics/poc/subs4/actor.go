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

type InboxCh chan Message
type DoneCh chan struct {}

type BasicActor struct {
	inbox InboxCh
}

func NewBasicActor() *BasicActor {
	return &BasicActor{
		inbox: make(InboxCh),
	}
}

func (actor *BasicActor) Inbox() InboxCh { return actor.inbox }

func (actor *BasicActor) Stop() {
	close(actor.Inbox())
}

func (actor *BasicActor) Start() {
	go func() {
		for mesg := range actor.Inbox() {
			switch mesg := mesg.(type) {
			case Command:
				mesg.Exec(actor)
			}
		}
	}()
}

func (actor *BasicActor) run() {

}