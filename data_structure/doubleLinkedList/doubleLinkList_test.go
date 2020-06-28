package doubleLinkedList

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	dlink := &DoubleLinkList{}
	dlink.Init(3)
	dlink.Append(1)
	dlink.Append("wocao")
	dlink.Append(3)
	fmt.Println(dlink.Full())
	dlink.Print()
	dlink.DeleteTail()
	dlink.Print()
	dlink.Add("?")
	// dlink.Add("*")

	dlink.Print()

	node := dlink.Get(1)

	dlink.MoveToHead(*node)

	dlink.Print()

}
