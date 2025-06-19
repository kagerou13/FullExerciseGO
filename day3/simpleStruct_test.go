package day3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	name string
	age  int
}

func (p Person) Greet() string {
	if p.age > 17 {
		return fmt.Sprintf("Welcome %s, your age is %d years old and you can enter this room", p.name, p.age)
	}
	return "You are under age"
}

func TestSimpleStruct(t *testing.T) {
	tests := []struct {
		name string
		req1 string
		req2 int
		exp  string
	}{
		{
			name: "first testing",
			req1: "Kage",
			req2: 20,
			exp:  "Welcome Kage, your age is 20 years old and you can enter this room",
		},
		{
			name: "second testing",
			req1: "Carmen",
			req2: 13,
			exp:  "You are under age",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p1 := Person{name: test.req1, age: test.req2}
			result := p1.Greet()
			assert.Equal(t, test.exp, result, "testing done")
		})
	}
}
