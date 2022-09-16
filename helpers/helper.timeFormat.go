package helpers

import (
	"time"
)

func TimeFormat(timeIso time.Time) string {
	return timeIso.String()[0:19]
}
