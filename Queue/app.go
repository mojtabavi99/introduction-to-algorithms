package main

func main() {
    queue := Queue{}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.enqueue(40)
	queue.enqueue(50)

	queue.dequeue()

	queue.print()
}
