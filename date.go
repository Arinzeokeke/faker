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

// Soon returns a soon time
func (d *Date) Soon(days int) time.Time {
	t := time.Now()
	tNano := t.Unix() * 1000

	tNano -= int64(randomIntRange(1000, days*24*3600*1000))

	return time.Unix(tNano/1000, 0)
}

// Month returns a  time month
func (d *Date) Month(abbr, context bool) string {
	var t string
	if abbr {
		t = "/abbr"
	} else {
		t = "/wide"
	}

	if context && directoryExists(d.DefaultLocale+"/"+datePrefix+t+"_context") {
		t += "_context"
	}

	return d.pick(datePrefix + t)
}
