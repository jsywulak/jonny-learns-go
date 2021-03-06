package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeouts(t *testing.T) {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		assert.True(t, false, "should not have gotten here")
		fmt.Println(res)
	case <-time.After(time.Second / 100):
		// fmt.Println("timeout 1")
		assert.True(t, true, "time out occured")
	}

	c2 := make(chan string, 1)

	go func() {
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		assert.Equal(t, "result 2", res)
	case <-time.After(time.Second):
		assert.True(t, false, "should not have gotten here")
	}
}
