package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	assert.Equal(t, 1, 1)

	match, _ := regexp.MatchString("p([[a-z]+])", "peach")
	fmt.Println(match)

}
