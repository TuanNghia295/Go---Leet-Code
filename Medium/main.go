package main

import "fmt"


type node struct{
	data int
	next *node
}

type linkedList struct{
	head *node
	length int
}


// method receiver (l *linkedList)
// With golang, define the method outside of the struct and this is how to indicated that

// the difference between method value receiver and pointer receiver is:
// We use pointer receiver when we want to work the original value
// we use value receiver when we just want to work with the coppy it

// This function use to add element at first index
func (l *linkedList) prepend (n *node)  {
	 second := l.head
	 l.head = n
	 l.head.next = second
	 l.length++
}


func (l linkedList) printListData()  {
	toPrint := l.head
	for  l.length != 0{
		fmt.Printf("%d ",toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Printf("\n")
}


func (l *linkedList) deleteWithValue(value int){
	// For example we have a linkedList [1,2,3,4]
	/*
	want to delete 3
	- compare value of each node, if next node = value want to delete, skip it
	*/

	// case delete with empty list
	if l.length == 0 {return}

	// case delete the first node
	if l.head.data == value{
		l.head = l.head.next
		l.length --
	}

	previousToDelete := l.head
	

	for previousToDelete.next.data != value{
		// case: delete non value in node list
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	// if equal with value, skip it
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main()  {
	myList := linkedList{}
	node1 := &node{data: 29}
	node2 := &node{data: 12}
	node3 := &node{data: 33}
	node4 := &node{data: 123}
	node5 := &node{data: 32}
	node6 := &node{data: 441}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.printListData()

	myList.deleteWithValue(111)
	myList.printListData()

}