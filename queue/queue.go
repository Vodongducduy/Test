package queue

const max = 100

type IQueue interface {
	Enqueue(val int)
	Dequeue()
	isEmpty() bool
}

type Queue struct {
	Value []int
	Front int
	Rear  int
}

func (q *Queue) Enqueue(val int) {
	q.Value = append(q.Value, val)
	if q.Front == -1 {
		q.Rear = 0
	}
	q.Front++
}

func (q *Queue) Dequeue() int {
	q.Front--
	if q.Front == -1 {
		return -1
	}
	deque := q.Value[q.Rear]
	q.Value = q.Value[q.Rear+1:]
	return deque
}

func (q *Queue) isEmpty() bool {
	if q.Front == -1 && q.Rear == -1 {
		return true
	}
	return false
}

func NewQueue() *Queue {
	return &Queue{
		Value: make([]int, 0, max),
		Front: -1,
		Rear:  -1,
	}
}
