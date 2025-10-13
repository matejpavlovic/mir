package priorityqueue

import "container/heap"

type WithPriority interface {
	Priority() int64
	Hash() []byte
}

type PriorityQueue[T WithPriority] struct {
	data queueData[T]
}

func New[T WithPriority]() *PriorityQueue[T] {
	pq := PriorityQueue[T]{
		data: queueData[T]{},
	}
	heap.Init(&(pq.data))
	return &pq
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.data.Len()
}

func (pq *PriorityQueue[T]) Push(value T) {
	heap.Push(&pq.data, value)
}

func (pq *PriorityQueue[T]) Pop() T {
	return heap.Pop(&pq.data).(T)
}
