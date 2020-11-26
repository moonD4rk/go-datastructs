package singly_linked_list

import "testing"

func TestNewListNode(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Range()
	m := l
	l.Insert(100, 5)
	l.Append(4)
	l.Len()
	l.Range()
	m.Range()
}
