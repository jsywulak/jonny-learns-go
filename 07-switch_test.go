package main

import (
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
		result := ""
		switch i.(type) {
		case bool:
			result = "boolean"
		case int:
			result = "integer"
		default:
			result = "unknown"
		}
		return result
	}

	assert.Equal(t, (whatAmI(true)), "boolean")
	assert.Equal(t, (whatAmI(1)), "integer")
	whatAmI(1)
	assert.Equal(t, (whatAmI("hey")), "unknown")
	whatAmI("hey")
}
