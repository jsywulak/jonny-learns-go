package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			assert.Equal(t, r, "a problem")
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
