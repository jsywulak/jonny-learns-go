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

	bob := person{"Bob", 20}
	assert.Equal(t, bob.name, "Bob")
	assert.Equal(t, bob.age, 20)
	alice := person{name: "Alice", age: 30}
	assert.Equal(t, alice.name, "Alice")
	assert.Equal(t, alice.age, 30)
	fred := person{name: "Fred"}
	assert.Equal(t, fred.name, "Fred")
	assert.Equal(t, fred.age, 0)
	ann := &person{name: "Ann", age: 40}
	assert.Equal(t, ann.name, "Ann")
	assert.Equal(t, ann.age, 40)
	jon := NewPerson("Jon")
	assert.Equal(t, jon.name, "Jon")
	assert.Equal(t, jon.age, 42)

	s := person{name: "Sean", age: 50}
	assert.Equal(t, s.name, "Sean")
	assert.Equal(t, s.age, 50)

	sp := &s
	assert.Equal(t, sp.age, 50)

	sp.age = 51
	assert.Equal(t, s.age, 51)
	fmt.Println(sp.age)
}
