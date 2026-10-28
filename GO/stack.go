package main

import "fmt"

type Stack[T any] struct {
	Head *Node[T]
}

func (s *Stack[T]) SEmpty() bool { return s.Head == nil }

func (s *Stack[T]) SPush(v T) {
	s.Head = &Node[T]{Value: v, Next: s.Head}
}

func (s *Stack[T]) SPop() {
	if s.SEmpty() {
		fmt.Println("ОШИБКА: Стек пуст, удалять нечего")
		return
	}
	s.Head = s.Head.Next
}

func (s *Stack[T]) STop() T {
	var zero T
	if s.SEmpty() {
		fmt.Println("ОШИБКА: Стек пуст")
		return zero
	}
	return s.Head.Value
}

func (s *Stack[T]) SPrint() {
	fmt.Print("S[")
	for c := s.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (s *Stack[T]) Clear() { s.Head = nil }
