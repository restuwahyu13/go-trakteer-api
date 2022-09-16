package helpers

import (
	"crypto/rand"
	"fmt"
)

func RandomToken() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
