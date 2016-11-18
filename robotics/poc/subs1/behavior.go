package subsumption
import (
	"github.com/fuzzycow/ev32go/helpers/monitor"
	"time"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/ev3api/device"
	"sync"
)

type PermitCh chan bool
type RecommendFunc func() int
type ActFunc func()

type Behavior interface {
	LeadsTo() Behavior
	Start()
	Act()
	Permit() chan bool
	Pause()
	Stop()
	Pri() int
}



type FuncBehavior struct {
	rec       RecommendFunc
	act       ActFunc
	interrupt chan struct {}
}

func NewFuncBehavior(rec RecommendFunc, act ActFunc) {
	&FuncBehavior{
		rec: rec,
		act: act,
	}
}

func (b *FuncBehavior) Start() {}

func (b *FuncBehavior) Next() Behavior {
	return b.rec()
}

func (b *FuncBehavior) Act() {
	return b.act()
}

type Arbiter struct {
	sb      [] Behavior
	permits PermitCh
	interval time.Duration
}

func NewArbiter(interval time.Duration, sb[] Behavior) {
	&Arbiter{
		sb: sb,
		interval: interval,
		permits: make(PermitCh),
	}
}

func (a *Arbiter) behaviorN(n int) {
	return a.sb[n]
}

func (a *Arbiter) Act() {
	var cur, next Behavior
	pollTicker := time.Ticker(a.interval)
	ACT: for {
		select {
		case permit, ok := <-a.permits:
			if !ok {
				break ACT
			}
		case <- pollTicker:
			if
			for i := 0; i < len(a.sb); i++ {
				b := a.sb[i]
				if cur != nil && b.Pri() <= cur {
					continue
				} else {
					next := b.LeadsTo()
					break
				}
			}

			if next == nil {
				continue
			}
			if cur != nil {
				cur.Pause()
			}
			next.Act()
		}
	}
}

type Permit interface {
	Suspend()
	Suspended()
	Revoked()
}

type LockPermit struct {
	mu sync.Mutex
}


func (b MontorSensor(sensor *device.Sensor) {
	for permit := range
}
