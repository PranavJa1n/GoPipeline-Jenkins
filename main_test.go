package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	a := 0
	for a != 2 {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(10) + 1
		result := hello(num)
		if result != "Hello" {
			t.Error("Expected Hello but got emply return")
			return
		}
		a += 1
	}
}
