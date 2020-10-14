package main

import "fmt"

// ArraySize is the size of the array for each bucket
const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each index of the hash table array
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hast table array
func (h *HashTable) Insert(key string) {
	index := hash(key, ArraySize)
	h.array[index].insert(key)
}

// Search will take in a key and return true if the key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key, ArraySize)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key, ArraySize)
	h.array[index].delete(key)
}

// insert will take in a key, create a node with the key and insert the node in the bocket
func (b *bucket) insert(k string) {
	// check if key is already in the bucket, if so, do nothing, keys are unique
	if b.search(k) {
		fmt.Printf("'%v' already exists\n", k)
		return
	}

	newNode := &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode

}

// search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

// delete takes in a key and delete the key to the bucket
func (b *bucket) delete(k string) {
	// if the current head is the key we are looking for,
	// reset the head to the next bucket in  the linked list
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	// we represent the current node as previous.next
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete by setting the next item in the previous node to the
			// what the next pointer of the node we want to delete points to
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

// hash transforms a string into an int by adding each rune value of each character
// in the string, and dividing it by the array size
func hash(key string, arrSize int) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % arrSize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hastTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		fmt.Printf("Hash value for %v is: %v\n", v, hash(v, ArraySize))
		hastTable.Insert(v)
	}

	hastTable.Delete("STAN")
	fmt.Println("STAN", hastTable.Search("STAN"))
	fmt.Println("BUTTERS", hastTable.Search("BUTTERS"))
}
