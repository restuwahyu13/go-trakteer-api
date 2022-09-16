package helpers

import (
	"time"
)

func ExpiredAt(timeDuration time.Duration, timeType string) int64 {
	switch timeType {
	case "months":
		return int64(time.Duration(time.Second*24*60*60*30) * timeDuration)

	case "days":
		return int64(time.Duration(time.Second*24*60*60) * timeDuration)

	case "hours":
		return int64(time.Duration(time.Hour) * timeDuration)

	case "minute":
		return int64(time.Duration(time.Minute) * timeDuration)

	case "second":
		return int64(time.Duration(time.Second) * timeDuration)

	case "milisecond":
		return int64(time.Duration(time.Millisecond) * timeDuration)

	default:
		return int64(time.Duration(time.Nanosecond) * timeDuration)
	}
}
