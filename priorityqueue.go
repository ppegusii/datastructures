package datastructures

import (
	"container/heap"
	"sync"
)

// NewPriorityQueue returns an implementation of the PriorityQueue interface thats
// built on top of http://golang.org/pkg/container/heap/.
// It features flexible priorities via user defined comparison and concurrency support.
func NewPriorityQueue(less Less, capacity int) PriorityQueue {
	// Possibly add ability to intialize with an array of items
	var pq *priorityQueue = &priorityQueue{
		itemHeap: make([]interface{}, 0, capacity),
		less:     less,
	}
	return pq
}

// Implements heap.Interface and Heap
type priorityQueue struct {
	sync.RWMutex
	itemHeap []interface{}
	less     Less
}

// Implementation of heap.Interface.
// Works with pqItems.

// Len is the number of elements in the collection.
func (this *priorityQueue) Len() int {
	return len(this.itemHeap)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (this *priorityQueue) Less(i, j int) bool {
	return this.less(this.itemHeap[i], this.itemHeap[j])
}

func (this *priorityQueue) Swap(i, j int) {
	var iItem interface{} = this.itemHeap[i]
	var jItem interface{} = this.itemHeap[j]
	this.itemHeap[j] = iItem
	this.itemHeap[i] = jItem
}

// add x as element Len()
func (this *priorityQueue) Push(x interface{}) {
	this.itemHeap = append(this.itemHeap, x)
}

// remove and return element Len() - 1.
func (this *priorityQueue) Pop() interface{} {
	if len(this.itemHeap) == 0 {
		return nil
	}
	var oldHeap []interface{} = this.itemHeap
	var oldHeapLen int = len(oldHeap)
	this.itemHeap = this.itemHeap[:oldHeapLen-1]
	// TODO shrink this.itemHeap underlying array if slice cap > 2 * len
	// possibly add variable that denotes prefer space over time efficiency
	return oldHeap[oldHeapLen-1]
}

// Implementation of PriorityQueue interface.
// Works with arbitrary items.

func (this *priorityQueue) Length() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.itemHeap)
}

func (this *priorityQueue) Head() interface{} {
	this.RLock()
	defer this.RUnlock()
	if this.Len() == 0 {
		return nil
	}
	return this.itemHeap[len(this.itemHeap)-1]
}

func (this *priorityQueue) Insert(x interface{}) {
	this.Lock()
	defer this.Unlock()
	heap.Push(this, x)
}

func (this *priorityQueue) PopTop() interface{} {
	this.Lock()
	defer this.Unlock()
	if this.Len() == 0 {
		return nil
	}
	return heap.Pop(this)
}
