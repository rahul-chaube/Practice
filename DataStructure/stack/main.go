package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	fmt.Println("**********  Welcome to stack ******* ")

	st := stack.Stack{}

	st.Push(10)
	st.Push(20)
	st.Push(30)

	fmt.Println(" ", st.Pop())
	fmt.Println(" ", st.Pop())
	fmt.Println("Is Empty ", st.IsEmpty())
	fmt.Println(" ", st.Pop())
	fmt.Println("Is Empty ", st.IsEmpty())

	fmt.Println("@@@@@@@@@@@@@@@@@@@@   Linked Stack @@@@@@@@@@@@@@")

	ls := stack.LinkedStack{}

	fmt.Println("  Size is ", ls.Size())

	ls.Push(10)
	ls.PrintAll()
	fmt.Println(" Print first ", ls.Peek())
	fmt.Println("  ", ls.Pop())
	fmt.Println("  ", ls.Pop())

	fmt.Println("@@@@@@@@ Queque @@@@@@@@@@  ")

	que := stack.Queqe{}
	fmt.Println(" Dequest ", que.DeQueue())

	que.EnQueue(10)
	que.EnQueue(20)

	que.PrintAll()

}
