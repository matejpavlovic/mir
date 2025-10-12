// This priority queue implementation is based on the example found at https://pkg.go.dev/container/heap

package priorityqueue

import (
	"bytes"
	"fmt"
)

type WithPriority interface {
	Priority() int
	Hash() []byte
}

// Item stored in the priority queue.
type Item[T WithPriority] struct {
	value T // The stored value.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items of type T.
type PriorityQueue[T WithPriority] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	pi, pj := pq[i].value.Priority(), pq[j].value.Priority()
	if pi != pj {
		return pi < pj
	} else {
		fmt.Printf("Comparing hashes: %s vs %s\n", pq[i].value.Hash(), pq[j].value.Hash())
		return bytes.Compare(pq[i].value.Hash(), pq[j].value.Hash()) < 0
	}
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := Item[T]{
		value: x.(T),
		index: n,
	}
	*pq = append(*pq, &item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item.value
}
