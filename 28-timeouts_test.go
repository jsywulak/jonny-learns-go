package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeouts(t *testing.T) {
	assert.Equal(t, 1, 1)
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		assert.True(t, false, "should not have gotten here")
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
		assert.True(t, true, "time out occured")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		assert.Equal(t, "result 2", res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
