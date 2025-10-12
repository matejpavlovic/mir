package priorityqueue

import (
	"container/heap"
	"github.com/stretchr/testify/require"
	"testing"
)

type fruit struct {
	priority int
	name     string
}

func (s *fruit) Priority() int {
	return s.priority
}

func (s *fruit) Hash() []byte {
	return []byte(s.name)
}

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue[*fruit], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item[*fruit]{
			&fruit{
				priority: priority,
				name:     value,
			},
			i,
		}
		i++
	}
	heap.Init(&pq)

	heap.Push(&pq, &fruit{priority: 1, name: "orange"})
	heap.Push(&pq, &fruit{priority: 8, name: "cherry"})
	heap.Push(&pq, &fruit{priority: 8, name: "grape"})
	heap.Push(&pq, &fruit{priority: 8, name: "apple"})

	require.Equal(t, "orange", heap.Pop(&pq).(*fruit).name)
	require.Equal(t, "apple", heap.Pop(&pq).(*fruit).name)
	require.Equal(t, "banana", heap.Pop(&pq).(*fruit).name)
	require.Equal(t, "pear", heap.Pop(&pq).(*fruit).name)
	require.Equal(t, "apple", heap.Pop(&pq).(*fruit).name)
	require.Equal(t, "cherry", heap.Pop(&pq).(*fruit).name)
	require.Equal(t, "grape", heap.Pop(&pq).(*fruit).name)
}
