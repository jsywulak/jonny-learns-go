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

}
