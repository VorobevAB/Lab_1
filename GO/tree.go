package main

import "fmt"

type CBNode[T any] struct {
	Value T
	Left  *CBNode[T]
	Right *CBNode[T]
}

type CompleteBinaryTree[T any] struct {
	Root *CBNode[T]
	Size int
}

func (t *CompleteBinaryTree[T]) TInsert(v T) {
	t.Root = t.insertAt(t.Root, t.Size, 0, v)
	t.Size++
}

func (t *CompleteBinaryTree[T]) insertAt(r *CBNode[T], target, idx int, v T) *CBNode[T] {
	if r == nil {
		if idx == target {
			return &CBNode[T]{Value: v}
		}
		return nil
	}
	if 2*idx+1 <= target {
		r.Left = t.insertAt(r.Left, target, 2*idx+1, v)
	}
	if 2*idx+2 <= target {
		r.Right = t.insertAt(r.Right, target, 2*idx+2, v)
	}
	return r
}

func (t *CompleteBinaryTree[T]) TFind(v T) {
	depth := 0
	if t.findRec(t.Root, v, &depth) {
		fmt.Printf("Элемент был найден на глубине %d\n", depth)
	} else {
		fmt.Println("Элемент не был найден в дереве")
	}
}

func (t *CompleteBinaryTree[T]) findRec(r *CBNode[T], v T, depth *int) bool {
	if r == nil {
		return false
	}
	if any(r.Value) == any(v) {
		return true
	}
	if t.findRec(r.Left, v, depth) {
		*depth++
		return true
	}
	if t.findRec(r.Right, v, depth) {
		*depth++
		return true
	}
	return false
}

func (t *CompleteBinaryTree[T]) TPrint() {
	fmt.Print("T[")
	t.inorder(t.Root)
	fmt.Println("]")
}

func (t *CompleteBinaryTree[T]) inorder(r *CBNode[T]) {
	if r == nil {
		return
	}
	t.inorder(r.Left)
	fmt.Printf("%v ", r.Value)
	t.inorder(r.Right)
}

func (t *CompleteBinaryTree[T]) TPrintTree() { t.printTree(t.Root, 0) }

func (t *CompleteBinaryTree[T]) printTree(r *CBNode[T], level int) {
	if r == nil {
		return
	}
	t.printTree(r.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	fmt.Println(r.Value)
	t.printTree(r.Left, level+1)
}

func (t *CompleteBinaryTree[T]) TIsComplete() bool {
	return t.isCompleteRec(t.Root, 0, t.Size)
}

func (t *CompleteBinaryTree[T]) isCompleteRec(r *CBNode[T], i, n int) bool {
	if r == nil {
		return true
	}
	if i >= n {
		return false
	}
	return t.isCompleteRec(r.Left, 2*i+1, n) &&
		t.isCompleteRec(r.Right, 2*i+2, n)
}

func (t *CompleteBinaryTree[T]) TIsFull() bool { return t.isFullRec(t.Root) }

func (t *CompleteBinaryTree[T]) isFullRec(r *CBNode[T]) bool {
	if r == nil {
		return true
	}
	if r.Left == nil && r.Right == nil {
		return true
	}
	if r.Left != nil && r.Right != nil {
		return t.isFullRec(r.Left) && t.isFullRec(r.Right)
	}
	return false
}

func (t *CompleteBinaryTree[T]) Clear() {
	t.Root = nil
	t.Size = 0
}
