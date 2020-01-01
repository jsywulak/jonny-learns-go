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
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
