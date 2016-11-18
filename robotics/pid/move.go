package pid

import (
	"time"
	"fmt"
)

type Move struct {
	Speed []int
	Duration  time.Duration
	RampUp time.Duration
	Fusion bool
	done chan struct{}
}

func NewMove() *Move {
	return &Move{
		done: make(chan struct{}),
	}
}

func (m *Move) WithSpeed(speed []int) *Move {
	m.Speed = speed
	return m
}

func (m *Move) WithRampUp(rampDuration time.Duration) *Move {
	m.RampUp = rampDuration
	return m
}

func (m *Move) WithDuration(d time.Duration) *Move {
	m.Duration = d
	return m
}

func (m *Move) WithFusion(fusion bool) *Move {
	m.Fusion = fusion
	return m
}

func (move Move) String() string {
	return fmt.Sprintf("speed=%+v,duration=%v, rampup=%v, fusion=%v", move.Speed, move.Duration,move.RampUp,move.Fusion)
}

func (m Move) Done() chan struct{} {
	return m.done
}

func (m *Move) Finish() {
	close(m.done)
}

