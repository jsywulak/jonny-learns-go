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

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Printf("I dunno type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
