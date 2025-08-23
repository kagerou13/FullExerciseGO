package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func deleteDuplicate(str string) string {
	seen := make(map[string]int)
	var newStr []string

	for _, v := range str {
		if v != 32 {
			seen[string(v)]++
		}
	}

	for _, value := range str {
		if elem, ok := seen[string(value)]; ok && elem == 1 {
			newStr = append(newStr, string(value))
		}
	}

	return strings.Join(newStr, "")
}

func TestDeleteDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		require  string
		expected string
	}{
		{
			name:     "first testing",
			require:  "kagerou valhalla",
			expected: "kgerouvh",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := deleteDuplicate(test.require)
			assert.Equal(t, test.expected, result, "testing done")
		})
	}
}
