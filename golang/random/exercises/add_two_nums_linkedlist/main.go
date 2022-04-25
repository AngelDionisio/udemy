package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Val  int
	Next *Node
}

func (n *Node) Print() {
	current := n
	for current != nil {
		fmt.Printf("value at current node: %d\n", current.Val)
		current = current.Next
	}
}

func (n *Node) GetListSize() int {
	counter := 0
	current := n
	for current != nil {
		counter++
		current = current.Next
	}
	return counter
}

func (n *Node) MakeList() []int {
	current := n
	list := make([]int, 0, n.GetListSize())
	for current != nil {
		list = append(list, current.Val)
		current = current.Next
	}
	return list
}

func getNum(head *Node) int {
	numList := head.MakeList()

	len := len(numList)
	var num int
	for i := len; i > 0; i-- {
		if i == len {
			num = numList[i-1]
			continue
		}
		num = (num * 10) + numList[i-1]
	}

	return num
}

func makeNode(val int, next *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
	}
}

func reverseInt(n int) int {
	result := 0
	for n > 0 {
		remainder := n % 10
		result *= 10
		result = result + remainder
		n /= 10
	}
	return result
}

func InsertNodeAtIthIndex(head *Node, index, data int) *Node {
	if head == nil {
		head = makeNode(data, nil)
		return head
	}
	if index == 0 {
		newNode := makeNode(data, nil)
		newNode.Next = head
		head = newNode
		return head
	}
	i := 0
	temp := head
	preNode := temp
	for temp != nil {
		if i == index {
			newNode := makeNode(data, nil)
			preNode.Next = newNode
			newNode.Next = temp
			break
		}
		i++
		preNode = temp
		temp = temp.Next
	}
	return head
}

func addTwoNumbers(list1 *Node, list2 *Node) *Node {
	num1 := getNum(list1)
	num2 := getNum(list2)

	sum := num1 + num2

	reversed := reverseInt(sum)
	fmt.Println("sum:", sum, "reversed:", reversed)
	reversedNumAsString := strconv.Itoa(reversed)

	var node *Node
	for i, v := range reversedNumAsString {
		converted, _ := strconv.Atoi(string(v))
		// fmt.Println("converted to int", converted)
		// node.Val = converted
		// // node.Next = makeNode()
		current := node
		newHead := InsertNodeAtIthIndex(current, i, converted)
		node = newHead
	}

	fmt.Println("node:", node.MakeList())
	return nil
}

func main() {
	linkedList1 := makeNode(2, makeNode(4, makeNode(3, nil)))
	linkedList2 := makeNode(5, makeNode(6, makeNode(4, nil)))

	fmt.Printf("list size: %d, values: %v\n", linkedList1.GetListSize(), linkedList1.MakeList())
	fmt.Printf("list size: %d, values: %v\n", linkedList2.GetListSize(), linkedList2.MakeList())

	fmt.Println("GetNum:", getNum(linkedList1))

	fmt.Println("addTwoNumbers:", addTwoNumbers(linkedList1, linkedList2))

	fmt.Println("about to exit")
}
