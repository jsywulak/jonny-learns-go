package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	name string
	age  int
}

func NewPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func TestStructs(t *testing.T) {

	bob := Person{"Bob", 20}
	assert.Equal(t, bob.name, "Bob")
	assert.Equal(t, bob.age, 20)
	alice := Person{name: "Alice", age: 30}
	assert.Equal(t, alice.name, "Alice")
	assert.Equal(t, alice.age, 30)
	fred := Person{name: "Fred"}
	assert.Equal(t, fred.name, "Fred")
	assert.Equal(t, fred.age, 0)
	ann := &Person{name: "Ann", age: 40}
	assert.Equal(t, ann.name, "Ann")
	assert.Equal(t, ann.age, 40)
	jon := NewPerson("Jon")
	assert.Equal(t, jon.name, "Jon")
	assert.Equal(t, jon.age, 42)

	s := Person{name: "Sean", age: 50}
	assert.Equal(t, s.name, "Sean")
	assert.Equal(t, s.age, 50)

	sp := &s
	assert.Equal(t, sp.age, 50)

	sp.age = 51
	assert.Equal(t, s.age, 51)
}
