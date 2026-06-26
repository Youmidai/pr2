package main

import (
	"fmt"
)

// СТЕК
type stack struct {
	s    []any // слайс, в котором хранятся значения
	head int   // индекс вершины стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в вершину стека
func push(s *stack, v any) {
	if s.head >= len(s.s)-1 {
		fmt.Println("Ошибка: стек переполнен!")
		return
	}
	s.head++
	s.s[s.head] = v
}

// pop - получение значения из стека и его удаление
func pop(s *stack) any {
	if s.head < 0 {
		return nil
	}
	val := s.s[s.head]
	s.s[s.head] = nil
	s.head--
	return val
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head < 0 {
		return nil
	}
	return s.s[s.head]
}

// БИНАРНОЕ ДЕРЕВО

type tree struct {
	head *node
}

type node struct {
	left, right *node
	v           int
}

func newTree() *tree {
	return &tree{}
}

// addTree - добавление значения в дерево
func addTree(t *tree, v int) {
	if t.head == nil {
		t.head = &node{v: v}
		return
	}
	insertNode(t.head, v)
}

func insertNode(n *node, v int) {
	if v < n.v {
		if n.left == nil {
			n.left = &node{v: v}
		} else {
			insertNode(n.left, v)
		}
	} else if v > n.v {
		if n.right == nil {
			n.right = &node{v: v}
		} else {
			insertNode(n.right, v)
		}
	}
}

// removeTree - удаление значения из дерева
func removeTree(t *tree, v int) {
	t.head = removeNode(t.head, v)
}

func removeNode(n *node, v int) *node {
	if n == nil {
		return nil
	}
	if v < n.v {
		n.left = removeNode(n.left, v)
	} else if v > n.v {
		n.right = removeNode(n.right, v)
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		minRight := n.right
		for minRight.left != nil {
			minRight = minRight.left
		}
		n.v = minRight.v
		n.right = removeNode(n.right, minRight.v)
	}
	return n
}

// valuesTree - получение отсортированного слайса
func valuesTree(t *tree) []int {
	var res []int
	inOrderTraverse(t.head, &res)
	return res
}

func inOrderTraverse(n *node, res *[]int) {
	if n == nil {
		return
	}
	inOrderTraverse(n.left, res)
	*res = append(*res, n.v)
	inOrderTraverse(n.right, res)
}

// ОДНОСВЯЗНЫЙ СПИСОК
type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

// addList - добавление значения в конец
func addList(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// get - получение по индексу
func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		return nil
	}
	curr := l.first
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr.v
}

// removeList - удаление по индексу
func removeList(l *singlyLinkedList, idx int) {
	if idx < 0 || idx >= l.size {
		return
	}
	if idx == 0 {
		l.first = l.first.next
		if l.first == nil {
			l.last = nil
		}
		l.size--
		return
	}

	prev := l.first
	for i := 0; i < idx-1; i++ {
		prev = prev.next
	}

	prev.next = prev.next.next
	if prev.next == nil {
		l.last = prev
	}
	l.size--
}

// valuesList - получение элементов в виде слайса
func valuesList(l *singlyLinkedList) []any {
	res := make([]any, 0, l.size)
	for curr := l.first; curr != nil; curr = curr.next {
		res = append(res, curr.v)
	}
	return res
}

func main() {
	s := newStack(2)
	push(s, "A")
	push(s, "B")
	fmt.Println("Стек pop:", pop(s))

	t := newTree()
	addTree(t, 2)
	addTree(t, 1)
	addTree(t, 3)
	fmt.Println("Дерево:", valuesTree(t))

	l := newSinglyLinkedList()
	addList(l, "X")
	addList(l, "Y")
	fmt.Println("Список:", get(l, 1))
}
