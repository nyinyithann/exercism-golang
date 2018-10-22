package clock

import (
	"fmt"
)

const (
	minutesPerHour = 60
	minutesPerDay  = 1440
)

// Clock represents hour and minute
type Clock struct{ h, m int }

// New instantiates clock
func New(h, m int) Clock {
	c := Clock{}
	r := (h*minutesPerHour + m) % minutesPerDay
	if r < 0 {
		r += minutesPerDay
	}
	c.h = r / minutesPerHour
	c.m = r % minutesPerHour
	return c
}

// Add increases given number of minutes
func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

// Subtract decreases given number of minutes
func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
