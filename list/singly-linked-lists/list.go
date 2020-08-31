package main

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
	if insertNum > l.Len() {
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

// Len get node list's length
func (l *List) Len() int {
	var count int
	if l == nil {
		return count
	}
	if l.Next == nil {
		return count + 1
	}
	for l.Next != nil {
		l = l.Next
		count++
	}
	return count + 1
}

// NewListNode create new list node
func NewListNode() *List {
	return &List{
		NodeNum:  "node0",
		NodeTime: time.Now(),
		Next:     nil,
	}
}

func main() {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Range()
	m := l
	l.Insert(100,5)
	l.Append(4)
	l.Len()
	l.Range()
	m.Range()
}

/*
output
node3 -> node100 -> node4 -> node5 -> node6 -> node7 -> node8 -> node9 ->
-----range is done
node0 -> node1 -> node2 -> node3 -> node100 -> node4 -> node5 -> node6 -> node7 -> node8 -> node9 ->
-----range is done
*/
