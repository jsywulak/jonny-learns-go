package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEpoch(t *testing.T) {
	assert.Equal(t, 1, 1)

	now := time.Date(
		2020, 01, 01, 01, 01, 01, 651387237, time.UTC)

	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000

	fmt.Println(now)
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
