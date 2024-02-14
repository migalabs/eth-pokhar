package utils

import (
	"math/rand"
	"time"
)

func GetRandomTimeout() time.Duration {
	// Generate a random timeout between 1 and 1.25 seconds
	return time.Duration(rand.Intn(250)+1000) * time.Millisecond
}
