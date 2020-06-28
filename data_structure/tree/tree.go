package tree

import (
	"math"
)

type ErChaTreeNode struct {
	Value int
	Left  *ErChaTreeNode
	Right *ErChaTreeNode
}

type BinarySearchTree struct {
	Root *ErChaTreeNode
}

func (t *BinarySearchTree) Insert(v int) {
	t.Root.Insert(v)
}

func (node *ErChaTreeNode) Insert(v int) {
	if v < node.Value {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &ErChaTreeNode{v, nil, nil}
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &ErChaTreeNode{v, nil, nil}
		}
	}
}

func (t *BinarySearchTree) InOrder() []int {
	var res []int
	t.Root.InOrder(&res)
	return res
}

func (node *ErChaTreeNode) InOrder(result *[]int) {
	if node.Left != nil {
		node.Left.InOrder(result)
	}
	*result = append(*result, node.Value)
	if node.Right != nil {
		node.Right.InOrder(result)
	}
}

func (node *ErChaTreeNode) PreOrder(result *[]int) {
	*result = append(*result, node.Value)
	if node.Left != nil {
		node.Left.PreOrder(result)
	}
	if node.Right != nil {
		node.Right.PreOrder(result)
	}
}

func (node *ErChaTreeNode) LastOrder(result *[]int) {
	if node.Left != nil {
		node.Left.LastOrder(result)
	}
	if node.Right != nil {
		node.Right.LastOrder(result)
	}
	*result = append(*result, node.Value)
}

func (node *ErChaTreeNode) GetDeep() int {
	if node == nil {
		return 0
	}
	lt, rt := node.Left, node.Right
	if lt.GetDeep() > rt.GetDeep() {
		return lt.GetDeep() + 1
	} else {
		return rt.GetDeep() + 1
	}
}

func (node *ErChaTreeNode) GetWidth() int {
	var lastWidth, currentWidth, maxWidth = 0, 0, 0
	query := QueueList{}
	if node != nil {
		query.Push(node)
		lastWidth = 1
	}
	for query.NotEmpty() {
		for lastWidth != 0 {
			lastWidth--
			node := query.GetHead()
			query.Pop()
			if node.Right != nil {
				query.Push(node.Right)
			}
			if node.Left != nil {
				query.Push(node.Left)
			}
		}
		currentWidth = query.Length()
		maxWidth = int(math.Max(float64(currentWidth), float64(maxWidth)))
		lastWidth = currentWidth
	}
	return maxWidth
}

type QueueList struct {
	list []*ErChaTreeNode
}

func (q *QueueList) Push(node *ErChaTreeNode) {
	q.list = append(q.list, node)
}

func (q *QueueList) NotEmpty() bool {
	if len(q.list) == 0 {
		return false
	}
	return true
}

func (q *QueueList) GetHead() *ErChaTreeNode {
	return q.list[0]
}

func (q *QueueList) Pop() {
	q.list = q.list[1:]
}

func (q *QueueList) Length() int {
	return len(q.list)
}
