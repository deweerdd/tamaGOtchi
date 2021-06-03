package main

import (
	"math/rand"
	"time"
)

func RandIdle() int {
	rand.Seed(time.Now().UnixNano())
	min := 6
	max := 14
	return rand.Intn(max-min+1) + min
}
