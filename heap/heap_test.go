package heap

import "testing"

func TestArrayHeap(t *testing.T) {
	heap := NewArrayHeap(20)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(50)
	heap.Insert(123)
	t.Log(heap.data)
}

func TestArrayHeap_BuildHeap(t *testing.T) {
	heap := NewArrayHeap(20)
	heap.BuildHeap([]int{20, 15, 50, 132, 87, 12, 65, 54, 53, 14, 22})
	t.Log(heap.data)
	heap.RemoveTop()
	t.Log(heap.data)
	heap.RemoveTop()
	t.Log(heap.data)
}
