package helpers

import (
	"regexp"
)

func ReplaceAllString(pattern, source, replace string) string {
	regex := regexp.MustCompile(pattern)
	res := regex.ReplaceAllString(source, replace)
	return res
}
