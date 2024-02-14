package utils

import (
	"math/rand"
	"time"
)

const (
	DefaultBaseTimeout = 1000 // in milliseconds
	DefaultTimeoutRand = 250  // in milliseconds

	ErrorCode429 = "429"
)

func GetRandomTimeout() time.Duration {
	// Generate a random timeout between 1 and 1.25 seconds
	return time.Duration(rand.Intn(DefaultTimeoutRand)+DefaultBaseTimeout) * time.Millisecond
}
