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
	assert.Equal(t, "2020-01-01 01:01:01.651387237 +0000 UTC", now.String())
	fmt.Println(secs)
	assert.Equal(t, int64(1577840461), secs)
	fmt.Println(millis)
	assert.Equal(t, int64(1577840461651), millis)
	fmt.Println(nanos)
	assert.Equal(t, int64(1577840461651387237), nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
