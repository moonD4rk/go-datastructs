package singly_linked_list

import (
	"fmt"
	"testing"
)

func TestNewListNode(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Range()
	m := l
	l.Insert(100, 5)
	l.Append(4)
	l.Length()
	l.Range()
	m.Range()
}

func TestList_Append(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Range()
}

func TestList_Length(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	len := l.Length()
	l.Range()
	fmt.Println(len)
}
