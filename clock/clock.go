package clock

type Clock struct {
	hours   int
	minutes int
}

func (c *Clock) Add(minutes int) {
	overallMinutes := c.minutes + minutes
	if overallMinutes >= 60 {
		overallHours := overallMinutes / 60
		overallHours = overallHours - ((overallHours / 24) * 24)
		overallMinutes = overallMinutes - (overallHours * 60)
		c.minutes = overallMinutes
		c.hours = overallHours
	}
	c.minutes = overallMinutes
}

func (c *Clock) Subtract(minutes int) {
	if minutes >= 60 {
		overallHours := minutes / 60
		overallHours = overallHours - ((overallHours / 24) * 24)
		overallMinutes := minutes - (overallHours * 60)
		c.minutes = overallMinutes
		if c.hours-overallHours < 0 {
			c.hours = 24 - (c.hours - overallHours)
		}
		c.hours = c.hours - overallHours

	}
	if c.hours == 0 && c.minutes-minutes < 0 {
		c.hours = 23
		c.minutes = (c.minutes - minutes) * (-1)
	}
}

func New(hours, minutes int) Clock {
	clock := Clock{hours, minutes}
	return clock
}
