package main

import "fmt"

// Node represents a single element in a doubly linked list.
// Each node contains an integer Value and pointers to the next and previous nodes.
type Node struct {
    Value int   // The value stored in the node
    Next  *Node // Pointer to the next node in the list (nil if this is the last node)
    Prev  *Node // Pointer to the previous node in the list (nil if this is the first node)
}

// LinkedList represents a doubly linked list.
type LinkedList struct {
    Head *Node // Pointer to the first node
    Tail *Node // Pointer to the last node
}

// Append adds a new node with the given value at the end of the list.
func (list *LinkedList) Append(value int) {
    node := &Node{Value: value}
    if list.Head == nil {
        list.Head = node
        list.Tail = node
        return
    }
    list.Tail.Next = node
    node.Prev = list.Tail
    list.Tail = node
}

// Prepend adds a new node with the given value at the beginning of the list.
func (list *LinkedList) Prepend(value int) {
    node := &Node{Value: value}
    if list.Head == nil {
        list.Head = node
        list.Tail = node
        return
    }
    node.Next = list.Head
    list.Head.Prev = node
    list.Head = node
}

// Length returns the number of nodes in the list.
func (list *LinkedList) Length() int {
    length := 0
    current := list.Head
    for current != nil {
        length++
        current = current.Next
    }
    return length
}

// Find searches for the first node with the given value.
func (list *LinkedList) Find(value int) *Node {
    current := list.Head
    for current != nil {
        if current.Value == value {
            return current
        }
        current = current.Next
    }
    return nil
}

// Get returns the node at the specified index (0-based).
// Returns false if index is out of bounds.
func (list *LinkedList) Get(index int) (*Node, bool) {
    if index < 0 || index >= list.Length() {
        return nil, false
    }

    current := list.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }

    return current, true
}

// Set updates the value of the node at the specified index.
func (list *LinkedList) Set(index int, value int) bool {
    node, ok := list.Get(index)
    if ok {
        node.Value = value
    }
    return ok
}

// Insert adds a new node with the given value at the specified index.
func (list *LinkedList) Insert(index int, value int) {
    if index < 0 || index > list.Length() {
        return
    }

    if index == 0 {
        list.Prepend(value)
        return
    }

    if index == list.Length() {
        list.Append(value)
        return
    }

    nextNode, _ := list.Get(index)
    prevNode := nextNode.Prev

    newNode := &Node{Value: value, Next: nextNode, Prev: prevNode}
    if prevNode != nil {
        prevNode.Next = newNode
    }
    nextNode.Prev = newNode
}

// Remove deletes the first node with the given value from the list.
func (list *LinkedList) Remove(value int) {
    current := list.Head
    for current != nil {
        if current.Value == value {
            if current.Prev != nil {
                current.Prev.Next = current.Next
            } else {
                list.Head = current.Next
            }
            if current.Next != nil {
                current.Next.Prev = current.Prev
            } else {
                list.Tail = current.Prev
            }
            return
        }
        current = current.Next
    }
}

// Print prints all values in the linked list.
func (list *LinkedList) Print() {
    current := list.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}
