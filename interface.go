package datastructures

/*
type ID uint

type Heap interface {
	// Re-establishes heap ordering after the element with
	// the given id changes its value in such a way that its priority changes.
	// Returns an error if the id is invalid.
	Fix(id ID) error
	Length() int
	// Return the item at the top of the heap along with its id.
	Peak() (interface{}, ID)
	// Returns the id of the pushed item, which can be used when calling Remove().
	PushItem(x interface{}) ID
	// Remove and return the item at the top of the heap along with its id.
	PopTop() (interface{}, ID)
	// Remove and return the item with the given id.
	// Returns an error if the id is invalid.
	Remove(id ID) (interface{}, error)
}
*/

// Function Less is the user defined comparison function.
// Less reports whether element x should sort before element y.
type Less func(x, y interface{}) bool

type PriorityQueue interface {
	Length() int
	Peak() interface{}
	PushItem(x interface{})
	PopTop() interface{}
}
