package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var p = fmt.Println

func TestStringFunctions(t *testing.T) {
	assert.Equal(t, 1, 1)

	p("Contains:", s.Contains("test", "es"))

}
