package singlylist

import (
	"fmt"
	"strings"
	"testing"
)

var (
	nodeValues = []int{1, 2, 3, 4, 5}
	sep        = "->"
)

func TestNewListNode(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  []int
		out string
	}{
		{
			in:  nodeValues,
			out: sliceToString([]int{0, 1, 2, 3, 4, 5}),
		},
	}
	l := New()
	for _, tc := range testCases {
		for _, value := range tc.in {
			l.Append(value)
		}
		out := l.Range(printNodeNum)
		if out != tc.out {
			t.Errorf("want %s, actual %s", tc.out, out)
		}
	}
}

func TestList_Insert(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  []int
		out string
	}{
		{
			in:  nodeValues,
			out: sliceToString([]int{0, 1, 2, 100, 3, 4, 5}),
		},
	}
	l := New()
	for _, tc := range testCases {
		for _, value := range tc.in {
			l.Append(value)
		}
		l.Insert(100, 3)
		out := l.Range(printNodeNum)
		if out != tc.out {
			t.Errorf("want %s, actual %s", tc.out, out)
		}
	}
}

func TestList_Append(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  []int
		out string
	}{
		{
			in:  nodeValues,
			out: sliceToString([]int{0, 1, 2, 3, 4, 5}),
		},
	}
	l := New()
	for _, tc := range testCases {
		for _, value := range tc.in {
			l.Append(value)
		}
		out := l.Range(printNodeNum)
		if out != tc.out {
			t.Errorf("want %s, actual %s", tc.out, out)
		}
	}
}

func TestList_Length(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  []int
		out int
	}{
		{
			in:  nodeValues,
			out: 6,
		},
	}
	l := New()
	for _, tc := range testCases {
		for _, value := range tc.in {
			l.Append(value)
		}
		out := l.Length()
		if out != tc.out {
			t.Errorf("want %d, actual %d", tc.out, out)
		}
	}
}

func TestList_RemoveNode(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  []int
		out string
	}{
		{
			in:  nodeValues,
			out: sliceToString([]int{0, 1, 2, 3, 5}),
		},
	}
	l := New()
	for _, tc := range testCases {
		for _, value := range tc.in {
			l.Append(value)
		}
		l.RemoveValue(3)
		out := l.Range(printNodeNum)
		if out != tc.out {
			t.Errorf("want %s, actual %s", tc.out, out)
		}
	}
}

func TestList_GetNode(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		in  []int
		out *List
	}{
		{
			in: nodeValues,
			out: &List{
				value: 3,
				next:  nil,
			},
		},
	}
	l := New()
	for _, tc := range testCases {
		for _, value := range tc.in {
			l.Append(value)
		}
		out := l.GetNode(3)
		if out.value != tc.out.value {
			t.Errorf("want %d, actual %d", tc.out.value, out.value)
		}
	}
}

func printNodeNum(value int) string {
	return fmt.Sprintf("%d%s", value, sep)
}

func sliceToString(s []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), sep), "[]") + sep
}
