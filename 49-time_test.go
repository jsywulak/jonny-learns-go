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

	fmt.Println(then)
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())

}
