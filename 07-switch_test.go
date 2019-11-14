package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSwitch(t *testing.T) {
	i := 2
	switch i {
	case 1:
		assert.True(t, false, "shouldn't get here")
	case 2:
		assert.True(t, true, "should've gotten here")
	case 3:
		assert.True(t, false, "shouldn't get here")
	}

	isWeekend := false
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		isWeekend = true
	default:
		isWeekend = false
	}

	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday {
		assert.True(t, isWeekend, "it's the weekend")
	} else {
		assert.False(t, isWeekend, "it's a weekday")
	}

	whatAmI := func(i interface{}) string {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("I dunno type %T\n", t)
		}
		return "hello"
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
