package shared

type Queue[T interface{}] struct {
	queue []*T
}

func NewQueue[T interface{}]() Queue[T] {
	return Queue[T]{
		queue: make([]*T, 0),
	}
}

func (q *Queue[T]) Offer(e *T) {
	q.queue = append(q.queue, e)
}

func (q *Queue[T]) Poll() (*T, bool) {
	if len(q.queue) == 0 {
		return nil, false
	}
	e := q.queue[0]
	q.queue = q.queue[1:]
	return e, true
}
