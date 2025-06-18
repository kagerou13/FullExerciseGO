package day2

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stack struct {
	element []int
}

func (s *Stack) Push(n int) {
	s.element = append(s.element, n)
}

func (s *Stack) Pop() int {
	if len(s.element) == 0 {
		os.Exit(1)
	}
	ln := len(s.element) - 1

	top := s.element[ln]       // last elemen
	s.element = s.element[:ln] // delele last elemen and update current element
	return top
}

func (s *Stack) Peek() int {
	if len(s.element) == 0 {
		os.Exit(0)
	}

	ln := len(s.element) - 1
	top := s.element[ln]

	return top
}

func TestStack(t *testing.T) {
	t.Run("testing Stack", func(t *testing.T) {
		s1 := Stack{}
		s1.Push(5)
		s1.Push(1412)
		s1.Pop()
		s1.Peek()
		assert.True(t, true)
	})
}

type Queue struct {
	elemen []int
}

func (q *Queue) Enqueue(n int) {
	q.elemen = append(q.elemen, n)
}

func (q *Queue) Dequeue() int {
	if len(q.elemen) == 0 {
		os.Exit(2)
	}

	return q.elemen[0]
}

func TestQueue(t *testing.T) {
	t.Run("testing queue", func(t *testing.T) {
		q1 := Queue{}
		q1.Enqueue(14)
		q1.Dequeue()
		assert.True(t, true)
	})
}
