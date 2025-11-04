package main

import (
	"fmt"
	"github.com/mojtabavi99/introduction-to-algorithms/pkg/structures"
)

func main() {
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}

	// Linked List Section
	fmt.Println("----------------------")
	fmt.Println("Linked List Section")
	fmt.Println("----------------------")

	linkedList := structures.LinkedList{}
	for _, v := range values {
		linkedList.Append(v)
	}

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

	stack := structures.Stack{}
	for _, v := range values {
		stack.Push(v)
	}

	fmt.Println("Stack height:", stack.Height)

	top, _ := stack.Peek()
	fmt.Println("Top element:", top)

	if !stack.IsEmpty() {
		stack.Pop()
	}

	stack.Print()

	// Queue Section
	fmt.Println("----------------------")
	fmt.Println("Queue Section")
	fmt.Println("----------------------")

	queue := structures.Queue{}
	for _, v := range values {
		queue.Enqueue(v)
	}

	fmt.Println("First element of Queue: ", queue.First.Value)
	fmt.Println("Last element of Queue: ", queue.Last.Value)
	fmt.Println("Length of Queue: ", queue.Length)

	fmt.Println("Dequeue executed")
	queue.Dequeue()

	queue.Print()

	// BST (binary search tree)
	fmt.Println("----------------------")
	fmt.Println("BST Section")
	fmt.Println("----------------------")

	bst := structures.BinarySearchTree{}
	for _, v := range values {
		bst.Insert(v)
	}

	bst.PrettyPrint()
	fmt.Println()

	fmt.Printf("In order: ")
	bst.InOrder()
	fmt.Println()

	fmt.Printf("Pre order: ")
	bst.PreOrder()
	fmt.Println()

	bst.Delete(40)
	fmt.Println("\nAfter delete 40")
	bst.PrettyPrint()

	// Graph
	fmt.Println("----------------------")
	fmt.Println("Graph Section")
	fmt.Println("----------------------")

	graph := structures.GraphInit(false)
	fmt.Println(graph)
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("A", "E")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "E")

	graph.PrettyPrint()
	fmt.Println("HasEdge(A, B):", graph.HasEdge("A", "B"))
	fmt.Println("HasEdge(C, E):", graph.HasEdge("C", "E"))

	fmt.Printf("Vertex D removed")
	graph.RemoveVertex("D")

	fmt.Println("\nAfter modifications:")
	graph.PrettyPrint()
}
