package main

import "fmt"

func main() {
    linkedList := LinkedList{}

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
}
