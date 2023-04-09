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

func (q *QueueRoute) SortAscending() {
	if q.nRoute > 1 {
		for i:= 1 ; i < q.nRoute; i++ {
			temp := q.buffer[i]
			j := i-1
			for temp.IsLessWeight(*q.buffer[j]) && j > 0 {
				q.buffer[j+1] = q.buffer [j]
				j--
			}
			if (j > 0){
				q.buffer[j] = temp
			} else {
				q.buffer[j+1] = q.buffer [j]
				q.buffer[j] = temp
			}
		}
	}
}
