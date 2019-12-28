package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var p = fmt.Println

func TestStringFunctions(t *testing.T) {
	assert.Equal(t, 1, 1)

	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ")
	fmt.Println("Repeat:    ")
	fmt.Println("Replace:   ")
	fmt.Println("Replace:   ")
	fmt.Println("Split:     ")
	fmt.Println("ToLower:   ")
	fmt.Println("ToUpper:   ")

}
