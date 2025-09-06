package main

import "fmt"

// LinkedList represents a doubly linked list
// with Head and Tail pointers for efficient access
// to the beginning and end of the list.
type LinkedList struct {
	Head *Node // Pointer to the first node in the list
	Tail *Node // Pointer to the last node in the list
}

// Append adds a new value at the end of the list.
// If the list is empty, the new node becomes both Head and Tail.
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

// Prepend adds a new value at the beginning of the list.
// If the list is empty, the new node becomes both Head and Tail.
func (list *LinkedList) Prepend(value int) {
	node := &Node{Value: value, Next: list.Head}
	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}

	node.Next = list.Head
	list.Head.Prev = node
	list.Head = node
}

// Delete removes the first node with the specified value from the list.
// If the node is Head or Tail, the corresponding pointers are updated.
func (list *LinkedList) Delete(value int) {
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

// Length returns the number of nodes in the linked list.
// It iterates from Head to Tail and counts each node.
func (list *LinkedList) Length() int {
    length := 0
    current := list.Head
    if current == nil {
        return 0
    }

    for current != nil {
        length++
        current = current.Next
    }

    return length
}


// Find searches for the first node with the specified value.
// Returns a pointer to the node if found, otherwise returns nil.
func (list *LinkedList) Find(value int) *Node {
	current := list.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
    fmt.Println("Node deleted ")
	return nil
}

// Print displays the list from Head to Tail.
func (list *LinkedList) Print() {
    if list.Head == nil {
        fmt.Println("Linked list is empty!")
    }
	current := list.Head
	for current != nil {
        if current != list.Tail {
            fmt.Printf("%d -> ", current.Value)
        } else {
            fmt.Println(list.Tail.Value)
        }
		current = current.Next
	}
}
