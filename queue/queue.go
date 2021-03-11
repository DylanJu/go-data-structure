package queue

type Queue struct {
	items []int
}

func (queue *Queue) Enqueue(value int) {
	queue.items = append(queue.items, value)
}

func (queue *Queue) Dequeue() int {
	first := queue.items[0]
	queue.items = queue.items[1:]
	return first
}

func MakeQueue() Queue {
	q := Queue{}

	return q
}
