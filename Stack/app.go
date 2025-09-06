package main

import "fmt"

func main() {
    stack := &Stack{}

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)
    stack.Push(40)

	fmt.Println("Stack Length:", stack.Length())

    top, _ := stack.Peek()
    fmt.Println("Top:", top)

    for !stack.IsEmpty() {
        value, _ := stack.Pop()
        fmt.Println("Pop:", value)
    }
}
