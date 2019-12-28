package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDefer(t *testing.T) {
	WrapperFunction(t)
	dat, err := ioutil.ReadFile("/tmp/dat")
	if err != nil {
		fmt.Println("you got bigger problems")
	}
	fmt.Println(dat)
}

func WrapperFunction(t *testing.T) {
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
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	err := f.Close()
	fmt.Fprintln(f, "closed")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
