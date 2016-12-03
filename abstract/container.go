// Package container abstract class defines a Push / Pop behavior
package container

// Iterator iterator interface
type Iterator interface {
	HasNext() bool
	Next() *Item
}

// Iterable iterable interface
type Iterable interface {
	Iterator() Iterator
}

// IContainer container interface
type IContainer interface {
	Iterable

	IsEmpty() bool
	Size() int
	Push(element interface{})
	Pop() *Item
}

// Item iterable item
type Item struct {
	Value interface{}
	Next  *Item
}

// container implements IContainer and Iterator
type container struct {
	first   *Item
	size    int
	current *Item
}

func (container *container) IsEmpty() bool {
	return container.first == nil
}

func (container *container) Size() int {
	return container.size
}

func (container *container) Iterator() Iterator {
	container.current = container.first
	return container
}

func (container *container) HasNext() bool {
	return container.current != nil
}

func (container *container) Next() *Item {
	item := container.current.Next
	container.current = item
	return item
}

/* abstract
func (container *container) Push(element interface{}) {
}

func (container *container) Pop(element interface{}) {
}
*/
