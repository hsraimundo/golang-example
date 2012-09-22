package main

import(
	"example/newmath"
	"gotest/singlelinkedlist"
	"fmt"
)

func main() {
	list := new(singlelinkedlist.SingleLinkedList)
	list.Add(nil)
	fmt.Printf("Hello, world. Sqrt(2) = %v...\n", newmath.Sqrt(2))
}