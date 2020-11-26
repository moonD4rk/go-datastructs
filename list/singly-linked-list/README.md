# 单向链表

**链表**是一种常见的基础数据结构，是一种线性表，但是并不会按线性的顺序存储数据，而是在每一个节点里存到下一个节点的指针(Pointer)。由于不必须按顺序存储，链表在插入的时候可以达到O(1)的复杂度，但是查找一个节点或者访问特定编号的节点则需要O(n)的时间。

使用链表结构可以克服数组链表需要预先知道数据大小的缺点，链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。但是链表失去了数组随机读取的优点，同时链表由于增加了结点的指针域，空间开销比较大。

链表中最简单的一种是**单向链表**，它包含两个域，一个信息域和一个指针域。这个链接指向列表中的下一个节点，而最后一个节点则指向一个空值。

一个单向链表的节点被分成两个部分。第一个部分保存或者显示关于节点的信息，第二个部分存储下一个节点的地址。单向链表只可向一个方向遍历。

链表最基本的结构是在每个节点保存数据和到下一个节点的地址，在最后一个节点保存一个特殊的结束标记，另外在一个固定的位置保存指向第一个节点的指针，有的时候也会同时储存指向最后一个节点的指针。一般查找一个节点的时候需要从第一个节点开始每次访问下一个节点，一直访问到需要的位置。但是也可以提前把一个节点的位置另外保存起来，然后直接访问。当然如果只是访问数据就没必要了，不如在链表上储存指向实际数据的指针。这样一般是为了访问链表中的下一个或者前一个（需要储存反向的指针，见下面的双向链表）节点。

相对于下面的双向链表，这种普通的，每个节点只有一个指针的链表也叫**单向链表**，或者**单链表**，通常用在每次都只会按顺序遍历这个链表的时候（例如图的邻接表，通常都是按固定顺序访问的）

![](https://darkmoon1973-cdn.oss-cn-beijing.aliyuncs.com/img/20201126135231.png)

### 初始化链表

```go
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
```



