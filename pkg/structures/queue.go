package structures

import "fmt"

// QueueNode represents a single element in a doubly linked list.
// Each node contains an integer Value and pointers to the next and previous nodes.
type QueueNode struct {
    Value int   // The value stored in the node
    Next  *QueueNode // Pointer to the next node in the list (nil if this is the last node)
}

// Queue represents a queue list.
type Queue struct {
    First *QueueNode // Pointer to the first node
    Last *QueueNode // Pointer to the last node
	Length int
}

func (queue *Queue) Enqueue(value int) {
	node := &QueueNode{Value: value}
	if  queue.Length == 0 {
		queue.First = node
		queue.Last = node
		queue.Length = 1
	} else {
		queue.Last.Next = node
		queue.Last = node
		queue.Length += 1
	}
}

func (queue *Queue) Dequeue() {
	if queue.Length != 0 {
		if queue.Length == 1 {
			queue.First = nil
			queue.Last = nil
			queue.Length = 0
		}

		temp := queue.First
		queue.First = temp.Next
		temp.Next = nil
		queue.Length -= 1
	}
}

func (queue *Queue) Print() {
	current_node := queue.First
    for current_node != nil {
        fmt.Printf("%d -> ", current_node.Value)
        current_node = current_node.Next
    }
    fmt.Println("nil")
}