package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DB struct {
	Arrays map[string]*Array[string]
	FLists map[string]*ForwardList[string]
	Lists  map[string]*List[string]
	Stacks map[string]*Stack[string]
	Queues map[string]*Queue[string]
	Trees  map[string]*CompleteBinaryTree[int]
}

func NewDB() *DB {
	return &DB{
		Arrays: make(map[string]*Array[string]),
		FLists: make(map[string]*ForwardList[string]),
		Lists:  make(map[string]*List[string]),
		Stacks: make(map[string]*Stack[string]),
		Queues: make(map[string]*Queue[string]),
		Trees:  make(map[string]*CompleteBinaryTree[int]),
	}
}

func split(s string) []string { return strings.Fields(s) }

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func exec(db *DB, line string) string {
	t := split(line)
	if len(t) == 0 || len(t[0]) == 0 {
		return "ОШИБКА: пустая команда"
	}
	
	op := t[0]
	
	if op == "PRINT" {
		if len(t) != 2 {
			return "ОШИБКА: PRINT <name>"
		}
		name := t[1]
		if db.Arrays[name] != nil {
			db.Arrays[name].MPrint()
			return "OK"
		}
		if db.FLists[name] != nil {
			db.FLists[name].FPrint()
			return "OK"
		}
		if db.Lists[name] != nil {
			db.Lists[name].LPrint()
			return "OK"
		}
		if db.Stacks[name] != nil {
			db.Stacks[name].SPrint()
			return "OK"
		}
		if db.Queues[name] != nil {
			db.Queues[name].QPrint()
			return "OK"
		}
		if db.Trees[name] != nil {
			db.Trees[name].TPrint()
			return "OK"
		}
		return "ОШИБКА: контейнер '" + name + "' не найден"
	}
	
	switch op[0] {
	case 'M':
		if op == "MPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: MPUSH <name> <value>"
			}
			if db.Arrays[t[1]] == nil {
				db.Arrays[t[1]] = &Array[string]{}
				db.Arrays[t[1]].Init()
			}
			db.Arrays[t[1]].MPush(t[2])
			return t[2]
		}
		if op == "MPUSHIDX" {
			if len(t) != 4 {
				return "ОШИБКА: MPUSHIDX <name> <idx> <value>"
			}
			if !isInt(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrays[t[1]] == nil {
				db.Arrays[t[1]] = &Array[string]{}
				db.Arrays[t[1]].Init()
			}
			idx, _ := strconv.Atoi(t[2])
			db.Arrays[t[1]].MPushIdx(idx, t[3])
			return t[3]
		}
		if op == "MGET" {
			if len(t) != 3 {
				return "ОШИБКА: MGET <name> <idx>"
			}
			if !isInt(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrays[t[1]] == nil {
				return ""
			}
			idx, _ := strconv.Atoi(t[2])
			return db.Arrays[t[1]].MGet(idx)
		}
		if op == "MDEL" {
			if len(t) != 3 {
				return "ОШИБКА: MDEL <name> <idx>"
			}
			if !isInt(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrays[t[1]] != nil {
				idx, _ := strconv.Atoi(t[2])
				db.Arrays[t[1]].MDelInd(idx)
			}
			return "OK"
		}
		if op == "MSET" {
			if len(t) != 4 {
				return "ОШИБКА: MSET <name> <idx> <value>"
			}
			if !isInt(t[2]) {
				return "ОШИБКА: idx — не число"
			}
			if db.Arrays[t[1]] == nil {
				db.Arrays[t[1]] = &Array[string]{}
				db.Arrays[t[1]].Init()
			}
			idx, _ := strconv.Atoi(t[2])
			db.Arrays[t[1]].MSetInd(idx, t[3])
			return t[3]
		}
		if op == "MLEN" {
			if len(t) != 2 {
				return "ОШИБКА: MLEN <name>"
			}
			if db.Arrays[t[1]] == nil {
				return "0"
			}
			return strconv.Itoa(db.Arrays[t[1]].MLen())
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'F':
		if op == "FPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: FPUSH <name> <value>"
			}
			if db.FLists[t[1]] == nil {
				db.FLists[t[1]] = &ForwardList[string]{}
			}
			db.FLists[t[1]].FPushFront(t[2])
			return t[2]
		}
		if op == "FPUSHB" {
			if len(t) != 3 {
				return "ОШИБКА: FPUSHB <name> <value>"
			}
			if db.FLists[t[1]] == nil {
				db.FLists[t[1]] = &ForwardList[string]{}
			}
			db.FLists[t[1]].FPushBack(t[2])
			return t[2]
		}
				if op == "FPUSHAFTER" {
			if len(t) != 4 {
				return "ОШИБКА: FPUSHAFTER <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *Node[string]
			for c := db.FLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.FLists[t[1]].FPushAfter(node, t[3])
			return t[3]
		}
		if op == "FPUSHBEFORE" {
			if len(t) != 4 {
				return "ОШИБКА: FPUSHBEFORE <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *Node[string]
			for c := db.FLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.FLists[t[1]].FPushBefore(node, t[3])
			return t[3]
		}
		if op == "FDELAFTER" {
			if len(t) != 3 {
				return "ОШИБКА: FDELAFTER <name> <ref>"
			}
			var node *Node[string]
			for c := db.FLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.FLists[t[1]].FDelAfter(node)
			return "OK"
		}
		if op == "FDELBEFORE" {
			if len(t) != 3 {
				return "ОШИБКА: FDELBEFORE <name> <ref>"
			}
			var node *Node[string]
			for c := db.FLists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.FLists[t[1]].FDelBefore(node)
			return "OK"
		}
		if op == "FDEL" {
			if len(t) != 2 {
				return "ОШИБКА: FDEL <name>"
			}
			if db.FLists[t[1]] != nil {
				db.FLists[t[1]].FDelFront()
			}
			return "OK"
		}
		if op == "FDELBACK" {
			if len(t) != 2 {
				return "ОШИБКА: FDELBACK <name>"
			}
			if db.FLists[t[1]] != nil {
				db.FLists[t[1]].FDelBack()
			}
			return "OK"
		}
		if op == "FDELVAL" {
			if len(t) != 3 {
				return "ОШИБКА: FDELVAL <name> <value>"
			}
			if db.FLists[t[1]] != nil {
				db.FLists[t[1]].FDelValue(t[2])
			}
			return "OK"
		}
		if op == "FFIND" {
			if len(t) != 3 {
				return "ОШИБКА: FFIND <name> <value>"
			}
			if db.FLists[t[1]] == nil {
				fmt.Println("Элемент не был найден в списке")
			} else {
				db.FLists[t[1]].FFind(t[2])
			}
			return "OK"
		}
		if op == "FPRINT" {
			if len(t) != 2 {
				return "ОШИБКА: FPRINT <name>"
			}
			if db.FLists[t[1]] == nil {
				fmt.Println("F[]")
			} else {
				db.FLists[t[1]].FPrint()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'L':
		if op == "LPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: LPUSH <name> <value>"
			}
			if db.Lists[t[1]] == nil {
				db.Lists[t[1]] = &List[string]{}
			}
			db.Lists[t[1]].LPushFront(t[2])
			return t[2]
		}
		if op == "LPUSHB" {
			if len(t) != 3 {
				return "ОШИБКА: LPUSHB <name> <value>"
			}
			if db.Lists[t[1]] == nil {
				db.Lists[t[1]] = &List[string]{}
			}
			db.Lists[t[1]].LPushBack(t[2])
			return t[2]
		}
				if op == "LPUSHAFTER" {
			if len(t) != 4 {
				return "ОШИБКА: LPUSHAFTER <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *LNode[string]
			for c := db.Lists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.Lists[t[1]].LPushAfter(node, t[3])
			return t[3]
		}
		if op == "LPUSHBEFORE" {
			if len(t) != 4 {
				return "ОШИБКА: LPUSHBEFORE <name> <ref> <value>"
			}
			if t[3] == "" {
				return "ОШИБКА: пустое значение"
			}
			var node *LNode[string]
			for c := db.Lists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.Lists[t[1]].LPushBefore(node, t[3])
			return t[3]
		}
		if op == "LDELAFTER" {
			if len(t) != 3 {
				return "ОШИБКА: LDELAFTER <name> <ref>"
			}
			var node *LNode[string]
			for c := db.Lists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.Lists[t[1]].LDelAfter(node)
			return "OK"
		}
		if op == "LDELBEFORE" {
			if len(t) != 3 {
				return "ОШИБКА: LDELBEFORE <name> <ref>"
			}
			var node *LNode[string]
			for c := db.Lists[t[1]].Head; c != nil; c = c.Next {
				if c.Value == t[2] {
					node = c
					break
				}
			}
			if node == nil {
				return "ОШИБКА: узла '" + t[2] + "' нет в списке"
			}
			db.Lists[t[1]].LDelBefore(node)
			return "OK"
		}
		if op == "LDEL" {
			if len(t) != 2 {
				return "ОШИБКА: LDEL <name>"
			}
			if db.Lists[t[1]] != nil {
				db.Lists[t[1]].LDelFront()
			}
			return "OK"
		}
		if op == "LDELBACK" {
			if len(t) != 2 {
				return "ОШИБКА: LDELBACK <name>"
			}
			if db.Lists[t[1]] != nil {
				db.Lists[t[1]].LDelBack()
			}
			return "OK"
		}
		if op == "LDELVAL" {
			if len(t) != 3 {
				return "ОШИБКА: LDELVAL <name> <value>"
			}
			if db.Lists[t[1]] != nil {
				db.Lists[t[1]].LDelValue(t[2])
			}
			return "OK"
		}
		if op == "LFIND" {
			if len(t) != 3 {
				return "ОШИБКА: LFIND <name> <value>"
			}
			if db.Lists[t[1]] == nil {
				fmt.Println("Элемент не был найден в списке")
			} else {
				db.Lists[t[1]].LFind(t[2])
			}
			return "OK"
		}
		if op == "LPRINT" {
			if len(t) != 2 {
				return "ОШИБКА: LPRINT <name>"
			}
			if db.Lists[t[1]] == nil {
				fmt.Println("L[]")
			} else {
				db.Lists[t[1]].LPrint()
			}
			return "OK"
		}
		if op == "LPRINTR" {
			if len(t) != 2 {
				return "ОШИБКА: LPRINTR <name>"
			}
			if db.Lists[t[1]] == nil {
				fmt.Println("Lr[]")
			} else {
				db.Lists[t[1]].LPrintReverse()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'S':
		if op == "SPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: SPUSH <name> <value>"
			}
			if db.Stacks[t[1]] == nil {
				db.Stacks[t[1]] = &Stack[string]{}
			}
			db.Stacks[t[1]].SPush(t[2])
			return t[2]
		}
		if op == "SPOP" {
			if len(t) != 2 {
				return "ОШИБКА: SPOP <name>"
			}
			if db.Stacks[t[1]] == nil || db.Stacks[t[1]].SEmpty() {
				return "ОШИБКА: стек пуст"
			}
			v := db.Stacks[t[1]].STop()
			db.Stacks[t[1]].SPop()
			return v
		}
		if op == "STOP" {
			if len(t) != 2 {
				return "ОШИБКА: STOP <name>"
			}
			if db.Stacks[t[1]] == nil || db.Stacks[t[1]].SEmpty() {
				return "ОШИБКА: стек пуст"
			}
			return db.Stacks[t[1]].STop()
		}
		if op == "SPRINT" {
			if len(t) != 2 {
				return "ОШИБКА: SPRINT <name>"
			}
			if db.Stacks[t[1]] == nil {
				fmt.Println("S[]")
			} else {
				db.Stacks[t[1]].SPrint()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'Q':
		if op == "QPUSH" {
			if len(t) != 3 {
				return "ОШИБКА: QPUSH <name> <value>"
			}
			if db.Queues[t[1]] == nil {
				db.Queues[t[1]] = &Queue[string]{}
			}
			db.Queues[t[1]].QPush(t[2])
			return t[2]
		}
		if op == "QPOP" {
			if len(t) != 2 {
				return "ОШИБКА: QPOP <name>"
			}
			if db.Queues[t[1]] == nil || db.Queues[t[1]].QEmpty() {
				return "ОШИБКА: очередь пуста"
			}
			v := db.Queues[t[1]].QFront()
			db.Queues[t[1]].QPop()
			return v
		}
		if op == "QGET" {
			if len(t) != 2 {
				return "ОШИБКА: QGET <name>"
			}
			if db.Queues[t[1]] == nil || db.Queues[t[1]].QEmpty() {
				return "ОШИБКА: очередь пуста"
			}
			return db.Queues[t[1]].QFront()
		}
		if op == "QPRINT" {
			if len(t) != 2 {
				return "ОШИБКА: QPRINT <name>"
			}
			if db.Queues[t[1]] == nil {
				fmt.Println("Q[]")
			} else {
				db.Queues[t[1]].QPrint()
			}
			return "OK"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	case 'T':
		if op == "TINSERT" {
			if len(t) != 3 {
				return "ОШИБКА: TINSERT <name> <int>"
			}
			if !isInt(t[2]) {
				return "ОШИБКА: значение — не число"
			}
			if db.Trees[t[1]] == nil {
				db.Trees[t[1]] = &CompleteBinaryTree[int]{}
			}
			v, _ := strconv.Atoi(t[2])
			db.Trees[t[1]].TInsert(v)
			return t[2]
		}
		if op == "TFIND" {
			if len(t) != 3 {
				return "ОШИБКА: TFIND <name> <int>"
			}
			if !isInt(t[2]) {
				return "ОШИБКА: значение — не число"
			}
			if db.Trees[t[1]] != nil {
				v, _ := strconv.Atoi(t[2])
				db.Trees[t[1]].TFind(v)
			}
			return "OK"
		}
		if op == "TPRINT" {
			if len(t) != 2 {
				return "ОШИБКА: TPRINT <name>"
			}
			if db.Trees[t[1]] == nil {
				fmt.Println("T[]")
			} else {
				db.Trees[t[1]].TPrint()
			}
			return "OK"
		}
		if op == "TPRINTTREE" {
			if len(t) != 2 {
				return "ОШИБКА: TPRINTTREE <name>"
			}
			if db.Trees[t[1]] != nil {
				db.Trees[t[1]].TPrintTree()
			}
			return "OK"
		}
		if op == "TISCOMPLETE" {
			if len(t) != 2 {
				return "ОШИБКА: TISCOMPLETE <name>"
			}
			if db.Trees[t[1]] != nil && db.Trees[t[1]].TIsComplete() {
				return "TRUE"
			}
			return "FALSE"
		}
		if op == "TISFULL" {
			if len(t) != 2 {
				return "ОШИБКА: TISFULL <name>"
			}
			if db.Trees[t[1]] != nil && db.Trees[t[1]].TIsFull() {
				return "TRUE"
			}
			return "FALSE"
		}
		return "ОШИБКА: неизвестная команда '" + op + "'"

	default:
		return "ОШИБКА: неизвестная команда '" + op + "'"
	}
}
