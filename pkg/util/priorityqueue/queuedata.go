package priorityqueue

import (
	"bytes"
)

// item stored in the priority queue.
type item[T WithPriority] struct {
	value T // The stored value.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A queueData implements heap.Interface and holds Items of type T.
type queueData[T WithPriority] []*item[T]

func (pq queueData[T]) Len() int { return len(pq) }

func (pq queueData[T]) Less(i, j int) bool {
	pi, pj := pq[i].value.Priority(), pq[j].value.Priority()
	if pi != pj {
		return pi < pj
	} else {
		return bytes.Compare(pq[i].value.Hash(), pq[j].value.Hash()) < 0
	}
}

func (pq queueData[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *queueData[T]) Push(x any) {
	n := len(*pq)
	item := item[T]{
		value: x.(T),
		index: n,
	}
	*pq = append(*pq, &item)
}

func (pq *queueData[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item.value
}
