package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEpoch(t *testing.T) {
	assert.Equal(t, 1, 1)

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000

	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
}
