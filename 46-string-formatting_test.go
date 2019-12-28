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

	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
