package singlelinkedlist

type SingleLinkedList struct{
	head, tail *SingleLinkedListNode
	length int
}

type SingleLinkedListNode struct{
	value string
	ref *SingleLinkedListNode
}

func (list *SingleLinkedList) Add(node *SingleLinkedListNode){
	
	if list.head == nil{
		list.head = node
		list.tail = node
	} else if list.head == list.tail{
		list.head.ref = node
		list.tail = node
	} else {
		list.tail.ref = node
		list.tail = node
	}

	list.length++

}

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


