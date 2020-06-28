package doubleLinkedList

import (
	"fmt"
	"sync"
)

type (
	DoubleLinkListNodeData interface{}

	DoubleLinkListNode struct {
		data DoubleLinkListNodeData
		pre  *DoubleLinkListNode
		next *DoubleLinkListNode
	}

	DoubleLinkList struct {
		lock *sync.Mutex
		cap  uint
		head *DoubleLinkListNode
		tail *DoubleLinkListNode
	}
)

func (l *DoubleLinkList) Init(size uint) {
	l.lock = new(sync.Mutex)
	l.cap = size
	l.head, l.tail = nil, nil
}

func (l *DoubleLinkList) Cap() uint {
	return l.cap
}

func (l *DoubleLinkList) Size() uint {
	length := uint(0)
	currNode := l.head
	if currNode == nil {
		return length
	}
	for currNode != nil {
		currNode = currNode.next
		length++
	}
	return length
}

func (l *DoubleLinkList) Full() bool {
	return l.cap <= l.Size()
}

func (l *DoubleLinkList) Get(idx int) *DoubleLinkListNode {
	var count int
	var currNode *DoubleLinkListNode

	currNode = l.head

	for count = 0; count < idx; count++ {
		currNode = currNode.next
	}
	return currNode
}

func (l *DoubleLinkList) GetHead() *DoubleLinkListNode {
	return l.head
}

func (l *DoubleLinkList) Add(data DoubleLinkListNodeData) {
	l.lock.Lock()
	defer l.lock.Unlock()
	newNode := &DoubleLinkListNode{
		data: data,
	}
	if l.Size() == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.pre = newNode
		l.head = newNode
		if l.Size() > l.Cap() {
			l.DeleteTail()
		}
	}

}

func (l *DoubleLinkList) Append(data DoubleLinkListNodeData) {
	l.lock.Lock()
	defer l.lock.Unlock()
	newNode := &DoubleLinkListNode{
		data: data,
	}
	if l.Size() >= l.Cap() {
		return
	}
	if l.Size() == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		if l.tail == nil {
			currNode := l.head
			for currNode.next != nil {
				currNode = currNode.next
			}
			l.tail = currNode
		}
		l.tail.next = newNode
		newNode.pre = l.tail
		l.tail = newNode
	}
}

func (l *DoubleLinkList) MoveToHead(node DoubleLinkListNode) {
	// 删除原有节点
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}
	node.next, node.pre = nil, nil
	// 添加新的队首节点
	l.Add(node.data)
}

func (l *DoubleLinkList) Print() {
	currNode := l.head
	for currNode != nil {
		fmt.Printf("%v ", currNode.data)
		currNode = currNode.next
	}
	fmt.Println("")
}

func (l *DoubleLinkList) DeleteTail() {
	l.tail = l.tail.pre
	l.tail.next.pre = nil
	l.tail.next = nil
}

func (l *DoubleLinkList) GetElements() []interface{} {
	var res = make([]interface{}, 0)
	var currNode = l.head
	for currNode != nil {
		res = append(res, currNode.data)
		currNode = currNode.next
	}
	return res
}
