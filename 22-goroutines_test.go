package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var dones = make(map[string]bool)
var mutex = &sync.Mutex{}

func f(from string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 1000)
	}
	mutex.Lock()
	dones[from] = true
	mutex.Unlock()
}

func TestRoutines(t *testing.T) {
	go f("goroutine")
	f("direct")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// just to make sure everything finished up
	time.Sleep(time.Second / 1000)
	assert.True(t, dones["goroutine"])
	assert.True(t, dones["direct"])

}
