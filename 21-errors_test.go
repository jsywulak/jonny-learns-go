package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func givemeanerror() error {
	return &argError{-1, "can't work with it"}
}

func TestErrors(t *testing.T) {
	assert.Equal(t, 1, 1)
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}
	r1p, e1 := f1(7)
	if assert.NoError(t, e1) {
		assert.Equal(t, 10, r1p)
	}
	r2, e2 := f1(42)
	if assert.Error(t, e2) {
		expected := errors.New("can't work with 42")
		assert.Equal(t, expected, e2)
	}

	fmt.Println(r2, e2)
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
