package main

import "fmt"

type Queue[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (q *Queue[T]) QEmpty() bool { return q.Head == nil }

func (q *Queue[T]) QPush(v T) {
	n := &Node[T]{Value: v}
	if q.Tail == nil {
		q.Head = n
		q.Tail = n
	} else {
		q.Tail.Next = n
		q.Tail = n
	}
}

func (q *Queue[T]) QPop() {
	if q.QEmpty() {
		fmt.Println("ОШИБКА: Очередь пуста, удалять нечего")
		return
	}
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
}

func (q *Queue[T]) QFront() T {
	var zero T
	if q.QEmpty() {
		fmt.Println("ОШИБКА: Очередь пуста")
		return zero
	}
	return q.Head.Value
}

func (q *Queue[T]) QPrint() {
	fmt.Print("Q[")
	for c := q.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (q *Queue[T]) Clear() { q.Head = nil; q.Tail = nil }
