package main

// if you need some explenations: https://ieftimov.com/golang-datastructures-linked-lists
import (
	"fmt"
)

type Node struct {
	Value int
	next  *Node
}

type List struct {
	head *Node
	size int
}

func (el *List) append(value int) {

	if el.size == 0 {
		el.head = &Node{value, nil}
	} else {
		currentPossition := el.head
		for currentPossition.next != nil {
			currentPossition = currentPossition.next
		}
		currentPossition.next = &Node{value, nil}
	}
	el.size++
}

func (el *List) print() {
	currentPossition := el.head
	for currentPossition != nil {
		fmt.Printf("%d ", currentPossition.Value)
		currentPossition = currentPossition.next
	}
}

func main() {

	fmt.Println("Gde si brejj")

	var LinkList List

	LinkList = List{}
	LinkList.append(1)
	LinkList.append(2)
	LinkList.append(3)
	LinkList.append(4)
	LinkList.append(5)

	// fmt.Printf("%d %d", LinkList.head.Value, LinkList.size)
	LinkList.print()

}
