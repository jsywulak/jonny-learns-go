package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRanges(t *testing.T) {
	assert.True(t, true)
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	assert.Equal(t, 9, sum)

	for i, num := range nums {
		if num == 3 {
			assert.Equal(t, 1, i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		assert.Equal(t, kvs[k], v)
	}

	for k := range kvs {
		assert.NotEmpty(t, kvs[k])
	}

	actual := []string
	for _, v := range kvs {
		fmt.Println("val:", v)
	}

	actual := ""
	for _, c := range "go" {
		actual += string(c)
	}
	assert.Equal(t, "go", actual)

	actual = ""
	for _, c := range "hello world" {
		actual += string(c)
	}
	assert.Equal(t, "hello world", actual)

}
