package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	assert.Equal(t, 2, 2)

	WrapperFunction(t)
}

func WrapperFunction(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	if true {
		panic("a problem")
	}

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
