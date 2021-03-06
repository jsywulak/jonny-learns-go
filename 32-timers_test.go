package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimers(t *testing.T) {
	timer1 := time.NewTimer(time.Second / 1000)
	<-timer1.C
	// fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		assert.Fail(t, "shouldn't execute this branch")
	}()
	stop2 := timer2.Stop()
	assert.NotNil(t, stop2, "stop 2 should be valued")

}
