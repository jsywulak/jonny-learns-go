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
	// for _, i := range []int{7, 42} {
	// 	if r, e := f1(i); e != nil {
	// 		fmt.Println("f1 failed:", e)
	// 	} else {
	// 		fmt.Println("f1 worked", r)
	// 	}
	// }
	r1p, e1p := f1(7)
	if assert.NoError(t, e1p) {
		assert.Equal(t, 10, r1p)
	}
	r1f, e1f := f1(42)
	assert.Equal(t, -1, r1f)
	if assert.Error(t, e1f) {
		expected := errors.New("can't work with 42")
		assert.Equal(t, expected, e1f)
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	r2p, e2p := f2(7)
	if assert.NoError(t, e2p) {
		assert.Equal(t, 10, r2p)
	}
	r2f, e2f := f2(42)
	assert.Equal(t, -1, r2f)
	if assert.Error(t, e2f) {
		// expected := errors.New("can't work with 42")
		// assert.Equal(t, expected, e2f)
		fmt.Println(r2f, e2f)
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		assert.Equal(t, 42, ae.arg)
		assert.Equal(t, "can't work with it", ae.prob)
	}
}
