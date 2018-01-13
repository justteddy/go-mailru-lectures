package main

import (
	"time"
)

// Calendar return struct with time.Time type
type Calendar struct {
	time.Time
}

// NewCalendar returns calendar
func NewCalendar(parsed time.Time) Calendar {
	return Calendar{parsed}
}

// CurrentQuarter returns quarter of the year by month num
func (n Calendar) CurrentQuarter() int {
	switch {
	case n.Month() <= 3:
		return 1
	case n.Month() <= 6:
		return 2
	case n.Month() <= 9:
		return 3
	}

	return 4
}
