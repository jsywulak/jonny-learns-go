package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefer(t *testing.T) {
	WrapperFunctionToBeAbleToDefer(t)
	dat, err := ioutil.ReadFile("/tmp/defer.txt")
	if err != nil {
		fmt.Println("you got bigger problems")
	}
	assert.Equal(t, "dataclosed", string(dat))
}

func WrapperFunctionToBeAbleToDefer(t *testing.T) {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Fprint(f, "data")
}

func closeFile(f *os.File) {
	fmt.Fprint(f, "closed")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
