package main

type Queue struct {
    items []int
}

func (q *Queue) Enqueue(value int) {
    q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
    if len(q.items) == 0 {
        return -1
    }

    front := q.items[0]
    q.items = q.items[1:]

    return front
}

func (q *Queue) Front() int {
    if len(q.items) == 0 {
        return -1
    }

    return q.items[0]
}

func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func (q *Queue) Size() int {
    return len(q.items)
}

func(q *Queue) Clear() {
	q.items = []int{}
}

func(q *Queue) Values() []int {
	values := make([]int, len(q.items))
	copy(values, q.items)
	return values
}

func(q *Queue) Items() []int {
	return q.items
}