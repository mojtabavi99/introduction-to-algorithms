package main

import "fmt"

func main() {
	linked_list := LinkedList{}

	linked_list.Append(20)
	linked_list.Append(30)
	linked_list.Append(40)

    linked_list.Prepend(10)

	node := linked_list.Find(30)
	if node != nil {
		fmt.Println(node.Value, " Founded")
	} else {
		fmt.Println("Node not found")
	}

    linked_list.Length()
    linked_list.Delete(20)
    linked_list.Length()

	linked_list.Print()
}
