package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	assert.Equal(t, 2, 2)

	if false {
		panic("a problem")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

func WrapperFunction() {

}
