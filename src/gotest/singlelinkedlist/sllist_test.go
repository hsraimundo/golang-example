package singlelinkedlist

import "testing"

func BuildList(even bool) *SingleLinkedList{
	var list = new(SingleLinkedList)
	node1 := SingleLinkedListNode {value: "First node"}
	node2 := SingleLinkedListNode {value: "Second node"}
	node3 := SingleLinkedListNode {value: "Third node"}
	node4 := SingleLinkedListNode {value: "Fourth node"}
	

	list.Add(&node1)
	list.Add(&node2)
	list.Add(&node3)
	if even {
		list.Add(&node4)
	}

	return list
}

func BuildEvenList() *SingleLinkedList{
	return BuildList(true)
}

func BuildOddList() *SingleLinkedList{
	return BuildList(false)
}

func TestFindMiddleEvenList_WithLength(t *testing.T){

	//Arrange
	var list = BuildEvenList()

	//Act
	first, second := list.FindMiddleWithLength()

	//Assert
	if first == nil || second == nil || first.value != "Second node" || second.value != "Third node" {
		t.Fail()
	}
}

func TestFindMiddleOddList_WithLength(t *testing.T){

	//Arrange
	var list = BuildOddList()

	//Act
	first, second := list.FindMiddleWithLength()

	//Assert
	if first == nil || second != nil || first.value != "Second node" {
		t.Fail()
	}
}

func TestFindMiddleEvenList_WithTwoPassages(t *testing.T){

	//Arrange
	var list = BuildEvenList()

	//Act
	first, second := list.FindMiddleWithTwoPassages()

	//Assert
	if first == nil || second == nil || first.value != "Second node" || second.value != "Third node" {
		t.Fail()
	}
}

func TestFindMiddleEvenList_WithTwoPointers(t *testing.T){

	//Arrange
	var list = BuildEvenList()

	//Act
	first, second := list.FindMiddleWithTwoPointers()

	//Assert
	if first == nil || second == nil || first.value != "Second node" || second.value != "Third node" {
		t.Fail()
	}
}

func TestFindMiddleOddList_WithTwoPassages(t *testing.T){

	//Arrange
	var list = BuildOddList()

	//Act
	first, second := list.FindMiddleWithTwoPassages()

	//Assert
	if first == nil || second != nil || first.value != "Second node" {
		t.Fail()
	}
}

func TestFindMiddleOddList_WithTwoPointers(t *testing.T){

	//Arrange
	var list = BuildOddList()

	//Act
	first, second := list.FindMiddleWithTwoPointers()

	//Assert
	if first == nil || second != nil || first.value != "Second node" {
		t.Fail()
	}
}

func BenchmarkFindMiddle(b *testing.B){
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


