package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func RandomDelay(min, max int) {
	delay := rand.Intn(max-min) + min
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
