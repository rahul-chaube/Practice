package stack

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedStack struct {
	head *Node
	size int
}

func (ls *LinkedStack) Size() int {
	return ls.size
}

func (ls *LinkedStack) Push(data int) {
	newNode := &Node{data: data}
	if ls.Size() == 0 {
		ls.head = newNode
	} else {
		newNode.next = ls.head
		ls.head = newNode
	}
	ls.size++
}
func (ls *LinkedStack) Peek() int {
	return ls.head.data
}

func (ls *LinkedStack) Pop() int {
	if ls.size == 0 {
		fmt.Println("Empty List ")
		return 0
	} else {
		data := ls.head.data
		ls.head = ls.head.next
		ls.size--
		return data
	}
}

func (ls *LinkedStack) PrintAll() {
	if ls.size == -1 {
		fmt.Println("Empty List ")

	} else {
		currentNode := ls.head

		for {
			fmt.Println("Data is ", currentNode.data)

			currentNode = currentNode.next

			if currentNode == nil {
				fmt.Println("End of list ")
				break
			}
		}
	}
}
