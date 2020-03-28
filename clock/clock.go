package clock

import (
	"fmt"
)

type Clock struct {
	hours   int
	minutes int
}

func (c *Clock) Add(minutes int) *Clock {
	overallMinutes := c.minutes + minutes
	if overallMinutes >= 60 {
		overallHours := overallMinutes / 60
		overallMinutes = overallMinutes - (overallHours * 60)
		overallHours = overallHours - ((overallHours / 24) * 24)
		c.minutes = overallMinutes
		if c.hours+overallHours >= 24 {
			c.hours = (24 - (c.hours + overallHours)) * (-1)
		} else {
			c.hours = c.hours + overallHours
		}
	} else {
		c.minutes = overallMinutes
	}
	return c
}

func (c *Clock) Subtract(minutes int) *Clock {
	overallMinutes := c.minutes - minutes
	if overallMinutes < 0 {
		c.hours = c.hours - 1
		overallHours := overallMinutes / 60
		overallMinutes = overallMinutes - (overallHours * 60)
		overallHours = overallHours - ((overallHours / 24) * 24)
		c.minutes = 60 + overallMinutes
		if c.hours-overallHours < 0 {
			c.hours = 24 - (c.hours - overallHours)
		}
		c.hours = c.hours - overallHours

	} else {
		c.minutes = c.minutes - minutes
	}
	if c.hours <= 0 && c.minutes-minutes < 0 {
		c.hours = 23
		c.minutes = (c.minutes - minutes) * (-1)
	}
	return c
}

func (c *Clock) String() (result string) {
	if c.hours < 10 {
		result = fmt.Sprintf("0%d:", c.hours)
	} else {
		result = fmt.Sprintf("%d:", c.hours)
	}
	if c.minutes < 10 {
		result = result + fmt.Sprintf("0%d", c.minutes)
	} else {
		result = result + fmt.Sprintf("%d", c.minutes)
	}
	return
}

func New(hours, minutes int) *Clock {
	hours, minutes = getHoursAndMinutes(hours, minutes)
	clock := Clock{hours, minutes}
	return &clock
}

func getHoursAndMinutes(hours, minutes int) (int, int) {
	if minutes >= 60 {
		additionalHours := minutes / 60
		minutes = minutes - (additionalHours * 60)
		hours = hours + additionalHours
	}
	if hours >= 24 {
		hours = hours - ((hours / 24) * 24)
	}
	return hours, minutes
}

func main() {
	clock := New(8, 0)
	fmt.Println(clock.String())
}
