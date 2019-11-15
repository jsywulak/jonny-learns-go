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
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println(kvs)

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for _, v := range kvs {
		fmt.Println("val:", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
	for i, c := range "hello world" {
		fmt.Println(i, c)
	}
}
