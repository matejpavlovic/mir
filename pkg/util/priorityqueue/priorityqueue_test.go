package priorityqueue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type fruit struct {
	priority int64
	name     string
}

func (s *fruit) Priority() int64 {
	return s.priority
}

func (s *fruit) Hash() []byte {
	return []byte(s.name)
}

func TestPriorityQueue(t *testing.T) {
	pq := New[*fruit]()

	pq.Push(&fruit{priority: 3, name: "banana"})
	pq.Push(&fruit{priority: 2, name: "apple"})
	pq.Push(&fruit{priority: 4, name: "pear"})
	pq.Push(&fruit{priority: 0, name: "orange"})
	pq.Push(&fruit{priority: 8, name: "cherry"})
	pq.Push(&fruit{priority: 8, name: "grape"})
	pq.Push(&fruit{priority: 8, name: "apple"})

	require.Equal(t, "orange", pq.Pop().name)
	require.Equal(t, "apple", pq.Pop().name)
	require.Equal(t, "banana", pq.Pop().name)
	require.Equal(t, "pear", pq.Pop().name)
	require.Equal(t, "apple", pq.Pop().name)
	require.Equal(t, "cherry", pq.Pop().name)
	require.Equal(t, "grape", pq.Pop().name)
}
