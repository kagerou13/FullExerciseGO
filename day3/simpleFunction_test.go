package day3

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(num1, num2 int) int {
	return num1 + num2
}

func subs(num1, num2 int) int {
	return num1 - num2
}

func multi(num1, num2 int) int {
	return num1 * num2
}

func div(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Divide by zero")
	}
	return num1 / num2, nil
}

func TestSample(t *testing.T) {
	tests := []struct {
		name       string
		req1, req2 int
		exp        int
		expZero    string
	}{
		{
			name: "add testing",
			req1: 15,
			req2: 13,
			exp:  28,
		},
		{
			name: "subs testing",
			req1: 28,
			req2: 13,
			exp:  15,
		},
		{
			name: "multi testing",
			req1: 14,
			req2: 4,
			exp:  56,
		},
		{
			name: "div testing",
			req1: 15,
			req2: 5,
			exp:  3,
		},
		{
			name:    "div testing",
			req1:    15,
			req2:    0,
			expZero: "Divide by zero",
		},
	}

	for _, test := range tests {
		switch test.name {
		case "add testing":
			t.Run(test.name, func(t *testing.T) {
				result := add(test.req1, test.req2)
				assert.Equal(t, test.exp, result, "add testing done")
			})
		case "subs testing":
			t.Run(test.name, func(t *testing.T) {
				result := subs(test.req1, test.req2)
				assert.Equal(t, test.exp, result, "subs testing done")
			})
		case "multi testing":
			t.Run(test.name, func(t *testing.T) {
				result := multi(test.req1, test.req2)
				assert.Equal(t, test.exp, result, "multi testing done")
			})
		case "div testing":
			t.Run(test.name, func(t *testing.T) {
				result, er := div(test.req1, test.req2)
				if er != nil {
					fmt.Println(er.Error())
				} else {
					assert.Equal(t, test.exp, result, "div testing done")
				}
			})

		}
	}
}
