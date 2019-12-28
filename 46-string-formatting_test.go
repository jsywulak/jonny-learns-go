package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type point struct {
	x, y int
}

func TestStringFormatting(t *testing.T) {
	assert.Equal(t, 1, 1)
	p := point{1, 2}
	result := fmt.Sprintf("%v", p)
	assert.Equal(t, "{1 2}", result)
	result = fmt.Sprintf("%+v", p)
	assert.Equal(t, "{x:1 y:2}", result)
	result = fmt.Sprintf("%#v", p)
	assert.Equal(t, "main.point{x:1, y:2}", result)
	result = fmt.Sprintf("%T", p)
	assert.Equal(t, "main.point", result)
	result = fmt.Sprintf("%d", 123)
	assert.Equal(t, "123", result)
	result = fmt.Sprintf("%b", 14)
	assert.Equal(t, "1110", result)
	result = fmt.Sprintf("%x", 456)
	assert.Equal(t, "1c8", result)
	result = fmt.Sprintf("%f", 78.9)
	assert.Equal(t, "78.900000", result)
	result = fmt.Sprintf("%e", 123400000.0)
	assert.Equal(t, "1.234000e+08", result)
	result = fmt.Sprintf("%E", 123400000.0)
	assert.Equal(t, "1.234000E+08", result)
	result = fmt.Sprintf("%s", "\"string\"")
	assert.Equal(t, "\"string\"", result)
	result = fmt.Sprintf("%q", "\"string\"")
	assert.Equal(t, "\"\\\"string\\\"\"", result)
	result = fmt.Sprintf("%x", "hex this")
	assert.Equal(t, "6865782074686973", result)
	fmt.Printf("%p", &p)
	fmt.Printf("|%6d|%6d|", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|", 1.2, 3.45)
	fmt.Printf("|%-6s|%-6s|", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s", "error")

}
