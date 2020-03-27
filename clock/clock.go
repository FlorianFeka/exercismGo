package clock

type Clock struct {
	hours   int
	minutes int
}

func (c *Clock) Add(minutes int) {
	overallMinutes := c.minutes + minutes
	if overallMinutes >= 60 {
		overallHours := overallMinutes / 60
		overallMinutes = overallMinutes - (overallHours * 60)
		for overallHours >= 24 {
			overallHours = overallHours - 24
		}
		c.minutes = overallMinutes
		c.hours = overallHours
	}
	c.minutes = overallMinutes
}

func New(hours, minutes int) Clock {
	clock := Clock{hours, minutes}
	return clock
}
