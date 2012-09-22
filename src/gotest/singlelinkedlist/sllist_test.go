package singlelinkedlist

import "testing"


//Base testing

func TestAdd(t *testing.T){
	//Arrange - Create an empty list and a node
	list := new(SingleLinkedList)
	node := SingleLinkedListNode {value: "node"}

	//Act - Add a node
	list.Add(&node)

	//Assert - length increment and references set
	if list.length != 1 {
		t.Fail()
	}
	if list.head != &node && list.tail != &node {
		t.Fail()
	}
}

func TestSingleNodeListString(t *testing.T){
	//Arrange - Create an empty list and a node
	list := new(SingleLinkedList)
	node := SingleLinkedListNode {value: "node"}

	//Act - Add a node
	list.Add(&node)

	//Assert
	if list.String() != "node" {
		t.Fail()
	}
}

func TestMultipleNodeListString(t *testing.T){
	//Arrange - Create an empty list and a node
	list := new(SingleLinkedList)
	node1 := SingleLinkedListNode {value: "node1"}
	node2 := SingleLinkedListNode {value: "node2"}

	//Act - Add a node
	list.Add(&node1)
	list.Add(&node2)

	//Assert
	if list.String() != "node1, node2" {
		t.Fail()
	}
}


//Test FindMiddleWithLength

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


//Test FindMiddleWithTwoPassages

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


//Test FindMiddleWithTwoPointers

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


/*
	Helper function to use on the Arrange fase of the tests.
	It creates a slingle linked list with even or odd nodes.
	even => even: true | odd: false
*/
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

//Build an even list.
func BuildEvenList() *SingleLinkedList{
	return BuildList(true)
}

//Build an odd list.
func BuildOddList() *SingleLinkedList{
	return BuildList(false)
}
