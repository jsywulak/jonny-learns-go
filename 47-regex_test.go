package main

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	assert.Equal(t, 1, 1)

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	assert.True(t, match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	assert.True(t, r.MatchString("peach"))

	assert.Equal(t, "peach", r.FindString("peach punch"))

	assert.Equal(t, []int{0, 5}, r.FindStringIndex("peach punch"))
	assert.Equal(t, []string{"peach", "ea"}, r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
