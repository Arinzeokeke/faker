package faker

import (
	"time"
)

// Dater Interface
type Dater interface {
	Past(years int) time.Time
	Future(years int) time.Time
	// Between
	// Recent
	// Soon
	// Month
	// Weekday
}

// Date struct
type Date struct {
	*Fake
}

// Past returns a past time
func (d *Date) Past(years int) time.Time {
	t := time.Now()
	tNano := t.Unix() * 1000

	tNano -= int64(randomIntRange(1000, years*365*24*3600*1000))

	return time.Unix(tNano/1000, 0)
}

// Future returns a future time
func (d *Date) Future(years int) time.Time {
	t := time.Now()
	tNano := t.Unix() * 1000

	tNano += int64(randomIntRange(1000, years*365*24*3600*1000))

	return time.Unix(tNano/1000, 0)
}
