package helpers

import (
	"time"
)

func TimeFormat(timeIso time.Time) string {
	return timeIso.Format("2006-01-02 15:04:05")
}
