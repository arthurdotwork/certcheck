package certcheck

import (
	"fmt"
	"math"
	"time"
)

type Certificate struct {
	CommonName string
	NotAfter   time.Time
}

func (c Certificate) FormatExpirationDate() string {
	until := c.NotAfter.Sub(time.Now())
	hours := math.Round(until.Hours())

	pluralize := func(s string, n int) string {
		if n == 1 {
			return s
		}
		return s + "s"
	}

	switch {
	case hours >= 24:
		return fmt.Sprintf("%d %s", int(hours/24), pluralize("day", int(hours/24)))
	case hours >= 1:
		return fmt.Sprintf("%d %s", int(hours), pluralize("hour", int(hours)))
	default:
		return fmt.Sprintf("%d %s", int(math.Round(until.Minutes())), pluralize("minute", int(math.Round(until.Minutes()))))
	}
}
