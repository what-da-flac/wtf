package timers

import "time"

type Timer struct{}

func New() *Timer {
	return &Timer{}
}

func (x *Timer) Now() time.Time {
	return time.Now()
}
