package fifo

import (
	"sync"
)

type FIFO[T any] interface {
	Add(item T)
	Next() (bool, T)
	Size() uint
	IsEmpty() bool
}

type node[T any] struct {
	value T
	next  *node[T]
}

type fifo[T any] struct {
	head  *node[T]
	tail  *node[T]
	mutex sync.Mutex
	size  uint
}

func New[T any]() FIFO[T] {
	return &fifo[T]{
		head:  nil,
		tail:  nil,
		mutex: sync.Mutex{},
		size:  0,
	}
}

func (f *fifo[T]) Add(item T) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	newNode := &node[T]{
		value: item,
		next:  nil,
	}

	if f.tail == nil {
		f.head = newNode
		f.tail = newNode
	} else {
		f.tail.next = newNode
		f.tail = newNode
	}

	f.size++
}

func (f *fifo[T]) Next() (bool, T) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.head == nil {
		var n T
		return false, n
	}

	item := f.head.value
	f.head = f.head.next
	if f.head == nil {
		f.tail = nil
	}

	if f.size > 0 {
		f.size--
	}

	return true, item
}

func (f *fifo[T]) Size() uint {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.size
}

func (f *fifo[T]) IsEmpty() bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return f.size == 0
}
