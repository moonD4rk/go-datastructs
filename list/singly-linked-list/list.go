package singly_linked_list

import (
	"fmt"
	"strconv"
	"time"
)

type List struct {
	NodeNum  interface{}
	NodeTime time.Time
	Next     *List
}

// NewListNode create new list node
func NewListNode() *List {
	return &List{
		NodeNum:  "node0",
		NodeTime: time.Now(),
		Next:     nil,
	}
}

// Append use for append one list node at tail
func (l *List) Append(nodeNum int) {
	newNode := &List{
		NodeNum:  "node" + strconv.Itoa(nodeNum),
		NodeTime: time.Now(),
		Next:     nil,
	}
	if l.Next == nil {
		l.Next = newNode
		return
	}
	// traversal linked list, find nil node
	for l.Next != nil {
		l = l.Next
	}
	l.Next = newNode
}

// Range use for traversal all list node
func (l *List) Range() {
	for l != nil {
		fmt.Printf("%v -> ", l.NodeNum)
		l = l.Next
	}
	fmt.Println()
	fmt.Printf("-----range is done\n")
}

// Insert use for insert node in specific num
func (l *List) Insert(nodeNum, insertNum int) {
	if insertNum > l.Length() {
		fmt.Println("insert > list length")
		return
	}
	newNode := &List{
		NodeNum:  "node" + strconv.Itoa(nodeNum),
		NodeTime: time.Now(),
		Next:     nil,
	}
	count := 1
	for l.Next != nil && count < insertNum {
		l = l.Next
		count++
	}
	newNode.Next = l.Next
	l.Next = newNode
}

// Len get length of list
func (l *List) Length() int {
	var count int
	// if list is nil, return 0
	if l == nil {
		return count
	}
	// if list have one value, return 1
	if l.Next == nil {
		return count + 1
	}
	for l.Next != nil {
		l = l.Next
		count++
	}
	return count + 1
}

/*
output
node3 -> node100 -> node4 -> node5 -> node6 -> node7 -> node8 -> node9 ->
-----range is done
node0 -> node1 -> node2 -> node3 -> node100 -> node4 -> node5 -> node6 -> node7 -> node8 -> node9 ->
-----range is done
*/
