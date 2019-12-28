package main

import (
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
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
