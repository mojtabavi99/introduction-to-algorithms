package main

import (
	"fmt"
	"github.com/mojtabavi99/introduction-to-algorithms/pkg/structures"
)

func main() {
	// Linked List Section
	fmt.Println("----------------------")
	fmt.Println("Linked List Section")
	fmt.Println("----------------------")
	linkedList := &structures.LinkedList{}

	linkedList.Append(20)
	linkedList.Append(30)
	linkedList.Append(40)
	linkedList.Append(50)
	linkedList.Append(60)
	linkedList.Prepend(10)

	fmt.Println("Initial list:")
	linkedList.Print()

	// Find node
	node := linkedList.Find(30)
	if node != nil {
		fmt.Println(node.Value, "found")
	} else {
		fmt.Println("Node not found")
	}

	// Get node at index
	node, result := linkedList.Get(2)
	if result {
		fmt.Println("Index 2:", node.Value)
	} else {
		fmt.Println("Index out of bounds")
	}

	// Set value
	linkedList.Set(4, 90)

	// Insert node
	linkedList.Insert(5, 100)

	fmt.Println("Final list:")
	linkedList.Print()

	// Stack Section
	fmt.Println("----------------------")
	fmt.Println("Stack Section")
	fmt.Println("----------------------")
	stack := &structures.Stack{}

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

	// Queue Section
	fmt.Println("----------------------")
	fmt.Println("Queue Section")
	fmt.Println("----------------------")
	queue := &structures.Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	queue.Dequeue()

	queue.Print()
}
