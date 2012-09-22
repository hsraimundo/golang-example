/*
	Package singlelinkedlist is a simple single linked list implementation in go.
	This implementations is for test purposes and pretends to demonstrate how to write in go,
	using references and pointers, and demonstrate go test features like:
		Unit tests;
		Benchmark tests.

	This package is defined in various files having the main implementation in the sllist.go file,
	the unit tests in the sllist_test.go and the benchmark tests in the sllist_bench_test.go file.

	For test purposes and having various cenarious for benchmarking this single linked list has a
	length counter, other than that it is based only on the sequence of list nodes with pointers
	between them.
*/
package singlelinkedlist

type SingleLinkedList struct{
	head, tail *SingleLinkedListNode
	length int
}

type SingleLinkedListNode struct{
	value string
	ref *SingleLinkedListNode
}

/*
	Add receives a new node reference to add to the current list.
	If the list is empty it's added to the head, otherwise it is added to the tail.

	The list's length is incremented using this method.
*/
func (list *SingleLinkedList) Add(node *SingleLinkedListNode){
	
	if list.head == nil{
		//if the list has no nodes, the head is the same as the tail
		list.head = node
		list.tail = node
	} else if list.head == list.tail{
		//if the list has one single node, updates the head reference
		// and adds a new node to the tail of the list.
		list.head.ref = node
		list.tail = node
	} else {
		//Adds a new node to the tail of the list
		list.tail.ref = node
		list.tail = node
	}

	//increment the list length
	list.length++

}

/*
	Finds the middle of the list based on the value of the list's length.
	If the length is even the two nodes that encircle the middle are returned,
	if it is odd the second return value is null and the middle node is returned in the 
	first return slot.
*/
func (list *SingleLinkedList) FindMiddleWithLength() (*SingleLinkedListNode, *SingleLinkedListNode){
	if list.length == 1 {
		return list.head, nil
	}
	var evenNumber = (list.length%2) == 0
	var middle int = 0
	middle = list.length/2

	node := list.head

	for i := 0; i<middle-1; i++{
		node = node.ref
	}

	if evenNumber {
		return node, node.ref
	}

	return node.ref, nil
}

/*
	Finds the middle of the list using a two passage technique. The first iteration through the list is used
	to obtain the list's length, following the same logic of the FindMiddleWithLength method afterwards.
*/
func (list *SingleLinkedList) FindMiddleWithTwoPassages() (*SingleLinkedListNode, *SingleLinkedListNode){
	pointer := list.head
	length := 0

	if pointer != nil {
		length++
	}

	for pointer.ref != nil {
		length++
		pointer = pointer.ref
	}

	//length of list now on 'length'
	pointer = list.head
	evenNumber := (length%2)==0

	var middle int = length/2

	for i := 0; i<middle-1; i++{
		pointer = pointer.ref
	}

	if evenNumber {
		return pointer, pointer.ref
	}

	return pointer.ref, nil

}

/*
	Finds the middle of the list using two iterations simultaneously. The first "iterator" goes node by node,
	and the second skips one node in each iteration.
	When the second pointer reaches the end of the list or passes the end (even or odd lists) it returns the middle
	nodes or node accordingly.
*/
func (list *SingleLinkedList) FindMiddleWithTwoPointers() (*SingleLinkedListNode, *SingleLinkedListNode){
	pointerA := list.head
	pointerB := list.head

	for pointerB.ref != nil && pointerB.ref.ref != nil {
		pointerA = pointerA.ref
		pointerB = pointerB.ref.ref
	}

	if pointerB.ref != nil && pointerB.ref.ref == nil {
		return pointerA, pointerA.ref
	}
	return pointerA, nil
}


/*
	Implementation of the Stringer interface.

	Interfaces in go are implemented by implementing the methos they contain. This way the
	interface implementation is implicit.

	This method returns a comma separated string with the values of the list.
*/
func (list *SingleLinkedList) String() string{
	ret := ""

	current := list.head

	ret += current.value
	for current != list.tail {
		current = current.ref
		ret+= ", " + current.value
	}

	return ret
}


