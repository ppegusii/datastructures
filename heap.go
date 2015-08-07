package datastructures

/*
import (
	"container/heap"
	"errors"
	"fmt"
	"log"
	"sync"
)

// http://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go

const mAXUINT = ^uint(0)
const mINUINT = 0
const mAXINT = int(mAXUINT >> 1)
const mININT = -mAXINT - 1

const (
	mIN ID = ID(mINUINT)
	mAX ID = ID(mAXUINT)
)

func NewHeap(less Less, capacity int) Heap {
	// Possibly add ability to intialize with an array of items
	var h *myHeap = &myHeap{
		itemHeap: make([]*hItem, 0, capacity),
		itemMap:  make(map[ID]*hItem),
		lastId:   mAX,
		less:     less,
	}
	return h
}

type hItem struct {
	id    ID
	index int
	value interface{}
}

func newHItem(id ID, index int, value interface{}) *hItem {
	return &hItem{
		id:    id,
		index: index,
		value: value,
	}
}

// Implements heap.Interface and Heap
type myHeap struct {
	sync.RWMutex
	itemHeap []*hItem
	itemMap  map[ID]*hItem
	less     Less
	lastId   ID
}

// Implementation of heap.Interface.
// Works with hItems.

// Len is the number of elements in the collection.
func (this *myHeap) Len() int {
	return len(this.itemHeap)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (this *myHeap) Less(i, j int) bool {
	return this.less(this.itemHeap[i], this.itemHeap[j])
}

func (this *myHeap) Swap(i, j int) {
	var iItem *hItem = this.itemHeap[i]
	var jItem *hItem = this.itemHeap[j]
	iItem.index = j
	this.itemHeap[j] = iItem
	jItem.index = i
	this.itemHeap[i] = jItem
}

// add x as element Len()
func (this *myHeap) Push(x interface{}) {
	i, _ := x.(*hItem)
	this.itemHeap = append(this.itemHeap, i)
	this.itemMap[i.id] = i
}

// remove and return element Len() - 1.
func (this *myHeap) Pop() interface{} {
	if len(this.itemHeap) == 0 {
		return nil
	}
	var oldHeap []*hItem = this.itemHeap
	var oldHeapLen int = len(oldHeap)
	this.itemHeap = this.itemHeap[:oldHeapLen-1]
	// TODO shrink this.itemHeap underlying array if slice cap > 2 * len
	// possibly add variable that denotes prefer space over time efficiency
	var i *hItem = oldHeap[oldHeapLen-1]
	delete(this.itemMap, i.id)
	return i
}

// Implementation of Heap interface.
// Works with arbitrary items.

func (this *myHeap) Fix(id ID) error {
	this.Lock()
	defer this.Unlock()
	var i *hItem
	var ok bool
	i, ok = this.itemMap[id]
	if !ok {
		return errors.New(fmt.Sprintf("Invalid id: %d", id))
	}
	heap.Fix(this, i.index)
	return nil
}

func (this *myHeap) Length() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.itemHeap)
}

func (this *myHeap) Peak() (interface{}, ID) {
	this.RLock()
	defer this.RUnlock()
	var length int = len(this.itemHeap)
	if length == 0 {
		return nil, 0
	}
	var i *hItem = this.itemHeap[len(this.itemHeap)-1]
	return i.value, i.id
}

func (this *myHeap) PushItem(x interface{}) ID {
	this.Lock()
	defer this.Unlock()
	var id ID = this.nextId()
	var i *hItem = newHItem(id, len(this.itemHeap), x)
	heap.Push(this, i)
	return id
}

func (this *myHeap) PopTop() (interface{}, ID) {
	this.Lock()
	defer this.Unlock()
	var i *hItem = heap.Pop(this).(*hItem)
	if i == nil {
		return nil, 0
	}
	return i.value, i.id
}

func (this *myHeap) Remove(id ID) (interface{}, error) {
	this.Lock()
	defer this.Unlock()
	var i *hItem
	var ok bool
	i, ok = this.itemMap[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Invalid id: %d", id))
	}
	heap.Remove(this, i.index)
	return i.value, nil
}

func (this *myHeap) nextId() ID {
	// TODO balanced binary tree of available ID regions
	// https://github.com/petar/GoLLRB
	var nextId ID = this.lastId
	if ID(len(this.itemMap)) == mAX {
		log.Fatal("Out of IDs")
	}
	for {
		if nextId != mAX {
			nextId++
		} else {
			nextId = mIN
		}
		_, ok := this.itemMap[nextId]
		if !ok {
			this.lastId = nextId
			return nextId
		}
	}
}
*/
