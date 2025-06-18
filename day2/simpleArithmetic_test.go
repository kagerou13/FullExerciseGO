package day2

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(num1, num2 float64) float64 {
	return num1 + num2
}

func substraction(num1, num2 float64) float64 {
	return num1 - num2
}

func div(num1, num2 float64) (float64, error) {
	if num2 < 1 {
		return 0.0, errors.New("Error: divide by zero")
	}
	return num1 / num2, nil
}

func multi(num1, num2 float64) float64 {
	return num1 * num2
}

func TestArithmetic(t *testing.T) {
	tests := []struct {
		name     string
		req1     float64
		req2     float64
		expected float64
	}{
		{
			name:     "addition testing",
			req1:     28,
			req2:     13,
			expected: 41,
		},
		{
			name:     "subtraction testing",
			req1:     28,
			req2:     13,
			expected: 15,
		},
		{
			name:     "multiplication testing",
			req1:     28,
			req2:     13,
			expected: 364,
		},
		{
			name:     "divide testing",
			req1:     28.0,
			req2:     13.0,
			expected: 2.1538461538461537,
		},
		{
			name:     "divide testing",
			req1:     28.0,
			req2:     0.0,
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := 0.0
			switch test.name {
			case "additon testing":
				result = add(test.req1, test.req2)
				assert.Equal(t, test.expected, result, "testing done")
			case "subtraction testing":
				result = substraction(test.req1, test.req2)
				assert.Equal(t, test.expected, result, "testing done")
			case "multiplication testing":
				result = multi(test.req1, test.req2)
				assert.Equal(t, test.expected, result, "testing done")
			case "divide testing":
				if result1, e := div(test.req1, test.req2); e != nil {
					fmt.Println(e)
				} else {
					result = result1
				}
				assert.Equal(t, test.expected, result, "testing done")
			}

		})

		// assert.Equal(t, test.expected, result, "testing done")
	}
}
