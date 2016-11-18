package monitor

import (
	"time"
)

type Policy int

const (
	Everything Policy = iota
	Changes
	Matches
)

type Monitor struct {
	stop chan struct{}
	interval time.Duration
	toTimer *time.Timer
	policy Policy
}

type PollFuncBool func() bool
type PollFuncInt func() int
type PollFuncString func() string

func New(policy Policy, interval time.Duration, timeout time.Duration) *Monitor {
	mon := &Monitor{
		stop: make(chan struct{}),
		interval: interval,
		toTimer: time.NewTimer(timeout),
		policy: policy,
	}

	return mon
}

func (mon *Monitor) PollBool(fn PollFuncBool) chan bool {
	ch := make(chan bool)
	go func(){
		defer close(ch)
		intervalTicker := time.NewTicker(mon.interval)
		defer intervalTicker.Stop()

		var lastRes bool

		POLLING: for {
			select {
			case <-intervalTicker.C:
				res := fn()
				switch {
				case mon.policy == Everything:
					ch <- res
				case mon.policy == Changes && res != lastRes:
					ch <- res
				case mon.policy == Matches && res:
					ch <- res
				}
				lastRes = res
			case <-mon.toTimer.C:
				break POLLING
			case <- mon.stop:
				break POLLING
			}
		}
	}()
	return ch
}

func (mon *Monitor) PollInt(fn PollFuncInt) chan int {
	ch := make(chan int)
	go func(){
		defer close(ch)
		intervalTicker := time.NewTicker(mon.interval)
		defer intervalTicker.Stop()

		var lastRes int

		POLLING: for {
			select {
			case <-intervalTicker.C:
				res := fn()
				switch {
				case mon.policy == Everything:
					ch <- res
				case mon.policy == Changes:
					if res != lastRes {
						ch <- res
					}
				default:
					panic("unsupported polling policy")
				}
				lastRes = res
			case <-mon.toTimer.C:
				break POLLING
			case <- mon.stop:
				break POLLING
			}
		}
	}()
	return ch
}

func (mon *Monitor) PollString(fn PollFuncString) chan string {
	ch := make(chan string)
	go func(){
		defer close(ch)
		intervalTicker := time.NewTicker(mon.interval)
		defer intervalTicker.Stop()

		var lastRes string

		POLLING: for {
			select {
			case <-intervalTicker.C:
				res := fn()
				switch {
				case mon.policy == Everything:
					ch <- res
				case mon.policy == Changes:
					if res != lastRes {
						ch <- res
					}
				default:
					panic("unsupported polling policy")
				}
				lastRes = res
			case <-mon.toTimer.C:
				break POLLING
			case <- mon.stop:
				break POLLING
			}
		}
	}()
	return ch
}



func (w *Monitor) Stop() {
	close(w.stop)
}

