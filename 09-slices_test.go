package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices(t *testing.T) {
	s := make([]string, 3)
	assert.ElementsMatch(t, s, [3]string{"", "", ""})

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	assert.ElementsMatch(t, s, [3]string{"a", "b", "c"})
	assert.Equal(t, "c", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	assert.ElementsMatch(t, s, [6]string{"a", "b", "c", "d", "e", "f"})

	c := make([]string, len(s))
	copy(c, s)
	assert.ElementsMatch(t, c, [6]string{"a", "b", "c", "d", "e", "f"})

	l := s[2:5]
	assert.ElementsMatch(t, l, [3]string{"c", "d", "e"})

	l = s[:5]
	assert.ElementsMatch(t, l, [5]string{"a", "b", "c", "d", "e"})

	l = s[2:]
	assert.ElementsMatch(t, l, [4]string{"c", "d", "e", "f"})

	t2 := []string{"g", "h", "i"}
	assert.ElementsMatch(t, t2, [3]string{"g", "h", "i"})

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	assert.ElementsMatch(t, [][]int{{0}, {1, 2}, {2, 3, 4}}, twoD)

}
