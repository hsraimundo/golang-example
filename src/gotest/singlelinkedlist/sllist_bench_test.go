package singlelinkedlist

import "testing"

func BenchmarkAdd(b *testing.B){
	
	for i:=0;i<b.N;i++{
		var list = new(SingleLinkedList)
		node1 := SingleLinkedListNode {value: "First node"}

		b.StartTimer()

		list.Add(&node1)

		b.StopTimer()
	}
}


func BenchmarkFindMiddleWithLength(b *testing.B){
	b.StopTimer()
	
	//Arrange
	var list = new(SingleLinkedList)
	node1 := SingleLinkedListNode {value: "First node"}
	node2 := SingleLinkedListNode {value: "Second node"}
	node3 := SingleLinkedListNode {value: "Third node"}
	node4 := SingleLinkedListNode {value: "Fourth node"}

	list.Add(&node1)
	list.Add(&node2)
	list.Add(&node3)
	list.Add(&node4)

	b.StartTimer()
	for i:=0;i<b.N;i++{
		//Act
		list.FindMiddleWithLength()
	}
}

func BenchmarkFindMiddleWithTwoPassages(b *testing.B){
	b.StopTimer()
	
	//Arrange
	var list = new(SingleLinkedList)
	node1 := SingleLinkedListNode {value: "First node"}
	node2 := SingleLinkedListNode {value: "Second node"}
	node3 := SingleLinkedListNode {value: "Third node"}
	node4 := SingleLinkedListNode {value: "Fourth node"}

	list.Add(&node1)
	list.Add(&node2)
	list.Add(&node3)
	list.Add(&node4)

	b.StartTimer()
	for i:=0;i<b.N;i++{
		//Act
		list.FindMiddleWithTwoPassages()
	}
}

func BenchmarkFindMiddleWithTwoPointers(b *testing.B){
	b.StopTimer()
	
	//Arrange
	var list = new(SingleLinkedList)
	node1 := SingleLinkedListNode {value: "First node"}
	node2 := SingleLinkedListNode {value: "Second node"}
	node3 := SingleLinkedListNode {value: "Third node"}
	node4 := SingleLinkedListNode {value: "Fourth node"}

	list.Add(&node1)
	list.Add(&node2)
	list.Add(&node3)
	list.Add(&node4)

	b.StartTimer()
	for i:=0;i<b.N;i++{
		//Act
		list.FindMiddleWithTwoPointers()
	}
}