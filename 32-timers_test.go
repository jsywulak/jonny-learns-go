package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimers(t *testing.T) {
	assert.Equal(t, 1, 1)

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

}
