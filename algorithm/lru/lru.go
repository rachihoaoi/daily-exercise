package lru

import (
	dList "github.com/rachihoaoi/daily-exercise.git/data_structure/doubleLinkedList"
)

type LRU struct {
	keyMap map[interface{}]*dList.DoubleLinkListNode
	list   dList.DoubleLinkList
}

func (lru *LRU) Init(cap uint) {
	lru.list = dList.DoubleLinkList{}
	lru.list.Init(cap)

	lru.keyMap = make(map[interface{}]*dList.DoubleLinkListNode, cap)
}

func (lru *LRU) Visit(data interface{}) {
	if lru.keyMap[data] == nil {
		lru.list.Add(data)
		lru.keyMap[data] = lru.list.GetHead()
	} else {
		node := lru.keyMap[data]
		lru.list.MoveToHead(*node)
	}
}

func (lru *LRU) Print() {
	lru.list.Print()
}

func (lru *LRU) GetElements() (res []interface{}) {
	return lru.list.GetElements()
}
