package main

import (
	"fmt"
)

// benefits of a linkedlist include O(1) insert time and space compelexity
// as you only need to create a new node, and the set points of the previous to new, and new to next of the previous node
// when inserting a value in a list for example, the algorithm would have to shift N number
// of items that follow to insert said item in the desired location
type Node struct {
	Val  string
	Next *Node
}

type IntNode struct {
	Val  int64
	Next *IntNode
}

func (n *Node) MakeListOfValues() []string {
	var l []string
	current := n
	for current != nil {
		l = append(l, current.Val)
		current = current.Next
	}
	return l
}

func makeNode(val string, next *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
	}
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		current = current.Next
	}
}

func PrintLinkedListRecursively(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("currentNodeValue: %s\n", head.Val)
	PrintLinkedList(head.Next)
}

func SumIntLinkedListValues(head *IntNode) int64 {
	current := head
	var sum int64
	for current != nil {
		sum += current.Val
		current = current.Next
	}
	return sum
}

func SumList(head *IntNode) int64 {
	if head == nil {
		return 0
	}
	return head.Val + SumList(head.Next)
}

/*
* if adding at the start of the linkedList, create a new node, and set the next value pointed to the old head
* to insert, create a variable to keep track of the "index" as you traverse the linkedList
* as long as the current node is not the tail, that is, it's not nil (it will traverse the list to nil, as the last one Next == nil)
 */
func InsertNode(head *Node, value string, index int) *Node {
	if index == 0 {
		newHead := makeNode(value, nil)
		newHead.Next = head
		return newHead
	}

	count := 0
	current := head
	for current != nil {
		if count == index-1 {
			next := current.Next
			newNode := makeNode(value, nil)
			current.Next = newNode
			newNode.Next = next
			// current.Next = makeNode(value, next)
		}
		count += 1
		current = current.Next
	}
	return head
}

func DeleteNode(head *Node, index int) *Node {
	if index == 0 {
		newHead := head.Next
		head.Next = nil
		return newHead
	}

	count := 0
	current := head
	for current != nil {
		if count == index-1 {
			current.Next = current.Next.Next
		}
		current = current.Next
		count++
	}
	return current
}

func ReverseLinkedList(head *Node) *Node {
	var prev *Node
	prev = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func FindTargetValue(head *Node, target string) bool {
	if head == nil {
		return false
	}
	current := head
	for current != nil {
		if current.Val == target {
			return true
		}
		current = current.Next
	}
	return false
}

func main() {
	a := makeNode("A", nil)
	b := makeNode("B", nil)
	c := makeNode("C", nil)
	d := makeNode("D", nil)
	a.Next = b
	b.Next = c
	c.Next = d

	// A -> B -> C -> D -> Nil
	PrintLinkedList(a)
	fmt.Println("list of Values:", a.MakeListOfValues())
	InsertNode(a, "1", 2)
	fmt.Println("list of Values after insert:", a.MakeListOfValues())
	DeleteNode(a, 2)
	fmt.Println("list of Values after deletion:", a.MakeListOfValues())

	reversed := ReverseLinkedList(a)
	fmt.Println("list of Values after reverse:", reversed.MakeListOfValues())

	target := "1"
	fmt.Printf("Does list contain value: %s, %v", target, FindTargetValue(a, target))
}
