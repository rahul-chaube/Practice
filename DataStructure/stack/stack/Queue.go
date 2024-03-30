package stack

import "fmt"

type Queqe struct {
	head *Node
	tail *Node
}

func (q *Queqe) EnQueue(data int) {
	newNode := &Node{data: data}

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queqe) DeQueue() int {

	if q.head == nil {
		fmt.Println("Queqe is empty")
		return 0
	} else {
		data := q.head.data
		q.head = q.head.next
		return data
	}
}

func (q *Queqe) PrintAll() {

	if q.head == nil {
		fmt.Println("Queqe is empty")

	} else {
		currentNode := q.head

		for {
			fmt.Println("Data is ", currentNode.data)

			currentNode = currentNode.next
			if currentNode == nil {
				return
			}
		}
	}
}
