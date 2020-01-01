package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	assert.Equal(t, 1, 1)

	now := time.Now()
	fmt.Println(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)

	assert.Equal(t, "2009-11-17 20:34:58.651387237 +0000 UTC", then.String())
	assert.Equal(t, 2009, then.Year())
	assert.Equal(t, "November", then.Month().String())
	assert.Equal(t, 17, then.Day())
	assert.Equal(t, 20, then.Hour())
	assert.Equal(t, 34, then.Minute())
	assert.Equal(t, 58, then.Second())
	assert.Equal(t, 651387237, then.Nanosecond())
	assert.Equal(t, "UTC", then.Location().String())
	assert.Equal(t, "Tuesday", then.Weekday().String())
	assert.True(t, then.Before(now))
	assert.False(t, then.After(now))
	assert.False(t, then.Equal(now))

	diff := now.Sub(then)
	fmt.Println(diff)

	assert.True(t, diff.Hours() > float64(88728.26597803133))
	assert.True(t, diff.Minutes() > float64(5.323697626748746e+06))
	assert.True(t, diff.Seconds() > float64(3.194218523991318e+08))
	assert.True(t, diff.Nanoseconds() > int64(319421847357361763))

	thenAgain := time.Date(
		2020, 01, 01, 01, 01, 01, 651387237, time.UTC)

	diff = thenAgain.Sub(then)

	fmt.Println(then.Add(diff))
	assert.Equal(t, "2020-01-01 01:01:01.651387237 +0000 UTC", then.Add(diff).String())
	fmt.Println(then.Add(-diff))
}
