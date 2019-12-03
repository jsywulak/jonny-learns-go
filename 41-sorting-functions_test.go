package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func TestSortingFunctions(t *testing.T) {
	assert.Equal(t, 1, 1)

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
