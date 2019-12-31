package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruit"`
}

func TestJson(t *testing.T) {
	assert.Equal(t, 1, 1)

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(1.23)
	fmt.Println(string(fltB))
}
