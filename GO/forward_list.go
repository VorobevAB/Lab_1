package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type ForwardList[T any] struct {
	Head *Node[T]
}

func (l *ForwardList[T]) FPushFront(v T) {
	l.Head = &Node[T]{Value: v, Next: l.Head}
}

func (l *ForwardList[T]) FPushBack(v T) {
	n := &Node[T]{Value: v}
	if l.Head == nil {
		l.Head = n
		return
	}
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
}

func (l *ForwardList[T]) FPushAfter(node *Node[T], v T) {
	if node == nil {
		l.FPushBack(v)
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = &Node[T]{Value: v, Next: cur.Next}
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *ForwardList[T]) FPushBefore(node *Node[T], v T) {
	if l.Head == node || node == nil {
		l.FPushFront(v)
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = &Node[T]{Value: v, Next: cur.Next}
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *ForwardList[T]) FDelFront() {
	if l.Head == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	l.Head = l.Head.Next
}

func (l *ForwardList[T]) FDelBack() {
	if l.Head == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	cur := l.Head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
}

func (l *ForwardList[T]) FDelAfter(node *Node[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur != nil {
		if cur.Next == nil {
			fmt.Printf("ОШИБКА: После узла %v ничего нет\n", node.Value)
			return
		}
		cur.Next = cur.Next.Next
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *ForwardList[T]) FDelBefore(node *Node[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if l.Head == node {
		fmt.Println("ОШИБКА: Перед головой ничего нет")
		return
	}
	if l.Head.Next == node {
		l.FDelFront()
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != nil && cur.Next.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = cur.Next.Next
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *ForwardList[T]) FDelNode(node *Node[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if l.Head == node {
		l.FDelFront()
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = node.Next
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *ForwardList[T]) FDelValue(v T) {
	cur := l.Head
	for cur != nil && any(cur.Value) != any(v) {
		cur = cur.Next
	}
	if cur != nil {
		l.FDelNode(cur)
	} else {
		fmt.Printf("ОШИБКА: Значения %v нет в списке\n", v)
	}
}

func (l *ForwardList[T]) FFind(v T) {
	i := 0
	for c := l.Head; c != nil; c = c.Next {
		if any(c.Value) == any(v) {
			fmt.Printf("Элемент был найден на позиции %d\n", i)
			return
		}
		i++
	}
	fmt.Println("Элемент не был найден в списке")
}

func (l *ForwardList[T]) FPrint() {
	fmt.Print("F[")
	for c := l.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (l *ForwardList[T]) FPrintRec(n *Node[T]) {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.Value)
	l.FPrintRec(n.Next)
}

func (l *ForwardList[T]) Clear() { l.Head = nil }
