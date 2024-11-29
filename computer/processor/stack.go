package processor

import (
	"fmt"
	"log/slog"
	"os"
)

type Stack struct {
	Items []interface{}
	size  int
	SP    int
}

func NewStack(size int) *Stack {
	return &Stack{
		Items: make([]interface{}, size),
		size:  size,
		SP:    0,
	}
}

func (s *Stack) Push(item interface{}) {
	if s.SP >= s.size {
		slog.Error("Push: Stack is full")
		os.Exit(1)
	}
	s.Items[s.SP] = item
	s.SP++
}

func (s *Stack) Pop() interface{} {
	if s.SP == 0 {
		slog.Error("Pop: Stack is empty")
		os.Exit(1)
	}
	s.SP--
	item := s.Items[s.SP]
	return item
}

func (s *Stack) Peek() interface{} {
	if s.SP == 0 {
		slog.Error("Peek: Stack is empty")
		os.Exit(1)
	}
	return s.Items[s.SP-1]
}

func (s *Stack) IsEmpty() bool {
	return s.SP == 0
}

func (s *Stack) String() string {
	str := ""
	for i := 0; i < s.SP; i++ {
		if s.Items[i] != nil {
			if i > 0 {
				str += " "
			}
			str += fmt.Sprintf("0x%x", s.Items[i])
		}
	}
	return "[" + str + "]"
}
