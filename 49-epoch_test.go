package main

import (
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

	assert.Equal(t, "2020-01-01 01:01:01.651387237 +0000 UTC", now.String())
	assert.Equal(t, int64(1577840461), secs)
	assert.Equal(t, int64(1577840461651), millis)
	assert.Equal(t, int64(1577840461651387237), nanos)

	assert.Equal(t, "2019-12-31 20:01:01 -0500 EST", time.Unix(secs, 0).String())
	assert.Equal(t, "2019-12-31 20:01:01.651387237 -0500 EST", time.Unix(0, nanos).String())
}
