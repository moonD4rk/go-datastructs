package singly_linked_list

import (
	"fmt"
	"time"
)

type List struct {
	NodeNum  int
	NodeTime time.Time
	Next     *List
}

// New create new list node
func New() *List {
	return &List{
		NodeNum:  0,
		NodeTime: time.Now(),
		Next:     nil,
	}
}

// Append use for append one list node at tail
func (l *List) Append(nodeNum int) {
	newNode := &List{
		NodeNum:  nodeNum,
		NodeTime: time.Now(),
		Next:     nil,
	}
	for iter := l; iter != nil; iter = iter.Next {
		if iter.Next == nil {
			iter.Next = newNode
			break
		}
	}
}

// Range use for traversal all list node
func (l *List) Range(rangFunc func(list *List)) {
	if l == nil {
		return
	}
	for iter := l; iter != nil; iter = iter.Next {
		rangFunc(iter)
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
		NodeNum:  nodeNum,
		NodeTime: time.Now(),
		Next:     nil,
	}
	count := 1
	for iter := l; iter != nil; iter = iter.Next {
		if count == insertNum {
			newNode.Next = iter.Next
			iter.Next = newNode
			break
		}
		count++
	}
}

func (l *List) RemoveNode(nodeNum int) {
	if l == nil {
		return
	}
	count := 0
	for iter := l; iter != nil; iter = iter.Next {
		if count == nodeNum {
			iter.Next = iter.Next.Next
			break
		}
		count++
	}
}

func (l *List) GetNode(nodeNum int) *List {
	if l == nil {
		return nil
	}
	count := 0
	for iter := l; iter != nil; iter = iter.Next {
		if count == nodeNum {
			h := &List{
				NodeNum:  iter.NodeNum,
				NodeTime: iter.NodeTime,
				Next:     nil,
			}
			return h
		}
		count++
	}
	return nil
}

// Length Len get length of list
func (l *List) Length() int {
	var count int
	// if list is nil, return 0
	if l == nil {
		return count
	}
	// if list have one value, return 1
	for iter := l; iter != nil; iter = iter.Next {
		count++
	}
	return count
}
