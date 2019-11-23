package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	assert.Equal(t, 1, 1)

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:", ints)
	ss := sort.StringsAreSorted(strs)
	fmt.Println("sorted?", ss)
	is := sort.IntsAreSorted(ints)
	fmt.Println("sorted?", is)
}
