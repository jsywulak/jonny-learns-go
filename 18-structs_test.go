package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func TestStructs(t *testing.T) {

	assert.Equal(t, 1, 1)
	bob := person{"Bob", 20}
	fmt.Println(bob)
	alice := person{name: "Alice", age: 30}
	fmt.Println(alice)
	fred := person{name: "Fred"}
	fmt.Println(fred)
	ann := &person{name: "Ann", age: 40}
	fmt.Println(ann)
	jon := NewPerson("Jon")
	fmt.Println(jon)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
