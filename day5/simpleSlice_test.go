package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var slice []int

func addElement(n int) {
	slice = append(slice, n)
}

func lenSlice() int {
	return len(slice)
}

func capSlice() int {
	return cap(slice)
}

type slice2 []int

func addElemen(slice *slice2, n int) {
	*slice = append(*slice, n)
}

func checkLenSlice(slice *slice2) int {
	return len(*slice)
}

func TestSimpleSlice(t *testing.T) {
	addElement(13)
	addElement(28)
	lenSlice()
	capSlice()

	var stack slice2
	addElemen(&stack, 3)
	checkLenSlice(&stack)

	assert.True(t, true)
}
