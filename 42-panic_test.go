package main

import (
	"fmt"
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
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
