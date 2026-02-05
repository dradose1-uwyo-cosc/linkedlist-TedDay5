package ds

import "errors"

type Queue struct {
	list LinkedList
}

// Push adds a tail node
func (q *Queue) Push(value string) {
	q.list.Insert(value)
}

// Pop removes the head node
func (q *Queue) Pop() (string, error) {
	if q.list.IsEmpty() {
		return "", errors.New("queue empty")
	}

	value := q.list.Head.data
	_ = q.list.RemoveAt(0)
	return value, nil
}
