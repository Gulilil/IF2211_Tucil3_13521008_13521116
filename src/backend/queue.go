package backend

type QueueRoute struct {
	buffer []*Route
	nRoute int
}

func (q *QueueRoute) Enqueue(r Route) {
	q.buffer = append(q.buffer, &r)
	q.nRoute++
}

func (q *QueueRoute) Dequeue() Route {
	temp := *q.buffer[0]
	q.buffer = q.buffer[1:]
	q.nRoute--
	return temp
}
