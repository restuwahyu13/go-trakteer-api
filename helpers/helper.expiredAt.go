package helpers

import (
	"time"
)

func ExpiredAt(timeDuration time.Duration, timeType string) time.Duration {
	switch timeType {
	case "months":
		return time.Duration(time.Second*24*60*60*30) * timeDuration

	case "days":
		return time.Duration(time.Second*24*60*60) * timeDuration

	case "hours":
		return time.Duration(time.Hour) * timeDuration

	case "minute":
		return time.Duration(time.Minute) * timeDuration

	case "second":
		return time.Duration(time.Second) * timeDuration

	case "milisecond":
		return time.Duration(time.Millisecond) * timeDuration

	default:
		return time.Duration(time.Nanosecond) * timeDuration
	}
}
