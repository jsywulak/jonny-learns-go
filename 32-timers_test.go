package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimers(t *testing.T) {
	assert.Equal(t, 1, 1)

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.println("timer 1 expired")

	timer2 := time.Newtimer(time.Second)
	go func() {
		<-timer2.C

	}()
}
