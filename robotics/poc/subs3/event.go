package main

type Event interface {
	Exec()
	Pri()
}

type EvenCh chan Event

type FnEvent struct {
	Fn func(Event) error
}

func NewFnEvent(fn func(Event)) Event {
	return &FnEvent{Fn: fn}
}

func (ev *FnEvent) Exec() {
	if ev.Fn != nil {
		ev.Fn()
	}
}

func main() {
	ev := NewFnEvent(func(Event) {})

	ev.Exec()
}
