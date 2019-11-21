package main

import (
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
}
