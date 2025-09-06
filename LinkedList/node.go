package main

// Node represents a single element in a doubly linked list.
// Each node contains an integer Value and pointers to the next and previous nodes.
type Node struct {
    Value int   // The value stored in the node
    Next  *Node // Pointer to the next node in the list (nil if this is the last node)
    Prev  *Node // Pointer to the previous node in the list (nil if this is the first node)
}
