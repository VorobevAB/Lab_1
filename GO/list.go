package main

import "fmt"

type LNode[T any] struct {
	Value T
	Next  *LNode[T]
	Prev  *LNode[T]
}

type List[T any] struct {
	Head *LNode[T]
	Tail *LNode[T]
}

func (l *List[T]) LPushFront(v T) {
	n := &LNode[T]{Value: v, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
}

func (l *List[T]) LPushBack(v T) {
	n := &LNode[T]{Value: v, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
}

func (l *List[T]) LPushBefore(node *LNode[T], v T) {
	if node == nil || l.Head == node {
		l.LPushFront(v)
		return
	}
	cur := l.Head
	for cur != nil && cur.Next != node {
		cur = cur.Next
	}
	if cur != nil {
		n := &LNode[T]{Value: v, Next: node, Prev: node.Prev}
		node.Prev.Next = n
		node.Prev = n
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *List[T]) LPushAfter(node *LNode[T], v T) {
	if node == nil {
		l.LPushBack(v)
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur != nil {
		n := &LNode[T]{Value: v, Next: node.Next, Prev: node}
		if node.Next != nil {
			node.Next.Prev = n
		} else {
			l.Tail = n
		}
		node.Next = n
	} else {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
	}
}

func (l *List[T]) LDelFront() {
	if l.Head == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	l.Head = l.Head.Next
	if l.Head != nil {
		l.Head.Prev = nil
	} else {
		l.Tail = nil
	}
}

func (l *List[T]) LDelBack() {
	if l.Tail == nil {
		fmt.Println("ОШИБКА: В списке нет элементов, удалять нечего")
		return
	}
	l.Tail = l.Tail.Prev
	if l.Tail != nil {
		l.Tail.Next = nil
	} else {
		l.Head = nil
	}
}

func (l *List[T]) LDelNode(node *LNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur == nil {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}
}

func (l *List[T]) LDelAfter(node *LNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if node == l.Tail {
		fmt.Println("ОШИБКА: За последним элементом, нет элементов для удаления")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur == nil {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
		return
	}
	if cur.Next == l.Tail {
		l.LDelBack()
	} else {
		cur.Next = cur.Next.Next
		cur.Next.Prev = cur
	}
}

func (l *List[T]) LDelBefore(node *LNode[T]) {
	if node == nil {
		fmt.Println("ОШИБКА: Узел не может быть nil")
		return
	}
	if node == l.Head {
		fmt.Println("ОШИБКА: Перед первым элементом, нет элементов для удаления")
		return
	}
	cur := l.Head
	for cur != nil && cur != node {
		cur = cur.Next
	}
	if cur == nil {
		fmt.Printf("ОШИБКА: Узла %v нет в списке\n", node.Value)
		return
	}
	if cur.Prev == l.Head {
		l.LDelFront()
	} else {
		cur.Prev = cur.Prev.Prev
		cur.Prev.Next = cur
	}
}

func (l *List[T]) LDelValue(v T) {
	cur := l.Head
	for cur != nil && any(cur.Value) != any(v) {
		cur = cur.Next
	}
	if cur != nil {
		l.LDelNode(cur)
	} else {
		fmt.Printf("ОШИБКА: Значения %v нет в списке\n", v)
	}
}

func (l *List[T]) LFind(v T) {
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

func (l *List[T]) LPrint() {
	fmt.Print("L[")
	for c := l.Head; c != nil; c = c.Next {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (l *List[T]) LPrintReverse() {
	fmt.Print("Lr[")
	for c := l.Tail; c != nil; c = c.Prev {
		fmt.Printf("%v ", c.Value)
	}
	fmt.Println("]")
}

func (l *List[T]) LPrintRec(n *LNode[T]) {
	if n == nil {
		return
	}
	fmt.Printf("%v ", n.Value)
	l.LPrintRec(n.Next)
}

func (l *List[T]) Clear() { l.Head = nil; l.Tail = nil }
