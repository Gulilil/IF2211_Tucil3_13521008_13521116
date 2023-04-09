package backend

type QueueRoute struct {
	buffer []*Route
	nRoute int
}

func (q *QueueRoute) Enqueue(r Route) {
	q.buffer = append(q.buffer, &r)
	q.nRoute++
}

func (q *QueueRoute) Dequeue() *Route {
	temp := *q.buffer[0]
	q.buffer = q.buffer[1:]
	q.nRoute--
	return &temp
}

func (q *QueueRoute) DequeueAt(idx int) Route {
	temp := *q.buffer[idx]
	for i := idx; i < q.nRoute-1; i++ {
		q.buffer[i] = q.buffer[i+1]
	}
	q.nRoute--
	return temp
}

func (q *QueueRoute) SortAscending() {
	if q.nRoute > 1 {
		for i := 1; i < q.nRoute; i++ {
			temp := &Route{}
			temp.CopyConstructorRoute(q.buffer[i])
			j := i - 1
			for temp.IsLessWeight(*q.buffer[j]) && j > 0 {
				q.buffer[j+1].CopyRoute(q.buffer[j])
				j--
			}
			if !temp.IsLessWeight(*q.buffer[j]) {
				q.buffer[j+1].CopyRoute(temp)
			} else {
				q.buffer[j+1].CopyRoute(q.buffer[j])
				q.buffer[j].CopyRoute(temp)
			}
		}
	}
}

func (q *QueueRoute) DisplayQueue() {
	for _, r := range q.buffer {
		r.DisplayRoute()
	}
}
