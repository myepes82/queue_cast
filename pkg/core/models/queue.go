package models

type Queue struct {
	items [][]byte
}

func NewQueue() *Queue {
	return &Queue{
		items: make([][]byte, 0),
	}
}

func (q *Queue) Enqueue(item []byte) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() []byte {
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *Queue) Peek() []byte {
	if len(q.items) == 0 {
		return nil
	}

	return q.items[0]
}
