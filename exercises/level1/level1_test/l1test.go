package level1test

import (
	"fmt"
)

type Node struct {
	Value string
}

func (n *Node) ToString() string {
	return fmt.Sprint(n.Value)
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

// Push: adds a new node to the stack
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop: removes and returns a node from the stack
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) ToString() string {
	stackStr := "data:{"
	for i := 0; i < len(s.nodes); i++ {
		stackStr += s.nodes[i].ToString() + ", "
	}
	stackStr += "}"
	return stackStr
}
