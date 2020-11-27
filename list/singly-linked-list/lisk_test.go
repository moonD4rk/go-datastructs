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
}

func TestList_Insert(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Range()
	l.Insert(100, 3)
	l.Range()
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
	fmt.Println(l.Length())
	if l.Length() != 5 {
		t.Errorf("length error")
	}

}

func TestList_RemoveNode(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Range()
	l.RemoveNode(3)
	l.Range()
}

func TestList_GetNode(t *testing.T) {
	l := NewListNode()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	r := l.GetNode(3)
	fmt.Println(r)
}
