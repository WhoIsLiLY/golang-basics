package main

import "fmt"

func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func main() {
	fmt.Println("Max int:", Max(10, 20))
	fmt.Println("Max string:", Max("apple", "banana"))

	stack := Stack[string]{}
	stack.Push("first")
	stack.Push("second")
	
	item, ok := stack.Pop()
	if ok {
		fmt.Println("Popped:", item)
	}
}