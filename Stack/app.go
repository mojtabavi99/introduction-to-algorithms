package main

import "fmt"

func main() {
    stack := &Stack{}

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)
    stack.Push(40)

	fmt.Println("Stack Height:", stack.Height)

    top, _ := stack.Peek()
    fmt.Println("Top:", top)

    if !stack.IsEmpty() {
        stack.Pop()
    }

    stack.Print()
}
