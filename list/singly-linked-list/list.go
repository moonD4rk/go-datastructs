package singlylist

import (
	"time"
)

type List struct {
	value      int
	createTime time.Time
	next       *List
}

// New create new list node
func New() *List {
	return &List{
		value:      0,
		createTime: time.Now(),
		next:       nil,
	}
}

// Append use for append one list node at tail
func (l *List) Append(value int) {
	newNode := &List{
		value:      value,
		createTime: time.Now(),
		next:       nil,
	}
	for iter := l; iter != nil; iter = iter.next {
		if iter.next == nil {
			iter.next = newNode
			break
		}
	}
}

// Range use for traversal all list node
func (l *List) Range(rangFunc func(value int) string) string {
	var s string
	if l == nil {
		return s
	}
	for iter := l; iter != nil; iter = iter.next {
		s += rangFunc(iter.value)
	}
	return s
}

// Insert use for insert node in specific num
func (l *List) Insert(nodeNum, insertNum int) {
	if insertNum > l.Length() {
		return
	}
	newNode := &List{
		value:      nodeNum,
		createTime: time.Now(),
		next:       nil,
	}
	count := 1
	for iter := l; iter != nil; iter = iter.next {
		if count == insertNum {
			newNode.next = iter.next
			iter.next = newNode
			break
		}
		count++
	}
}

func (l *List) RemoveValue(value int) {
	if l == nil {
		return
	}
	count := 0
	for iter := l; iter != nil; iter = iter.next {
		if count == value {
			iter.next = iter.next.next
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
	for iter := l; iter != nil; iter = iter.next {
		if count == nodeNum {
			h := &List{
				value:      iter.value,
				createTime: iter.createTime,
				next:       nil,
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
	for iter := l; iter != nil; iter = iter.next {
		count++
	}
	return count
}
