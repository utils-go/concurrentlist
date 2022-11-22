package concurrentlist

import (
	"reflect"
	"sync"
)

type ConcurrentListT[T any] struct {
	data []T
	mux  sync.Mutex
}

func NewList[T any]() *ConcurrentListT[T] {
	return &ConcurrentListT[T]{
		data: make([]T, 0),
	}
}

func (c *ConcurrentListT[T]) Add(v T) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = append(c.data, v)
}

func (c *ConcurrentListT[T]) AddRange(v []T) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = append(c.data, v...)
}

func (c *ConcurrentListT[T]) Clear() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = make([]T, 0)
}

func (c *ConcurrentListT[T]) Remove(v T) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	newslice := make([]T, 0)
	isexist := false

	for _, d := range c.data {
		if reflect.DeepEqual(d, v) && !isexist {
			isexist = true
		} else {
			newslice = append(newslice, d)
		}
	}

	if isexist {
		c.data = newslice
		return true
	}

	return false
}

func (c *ConcurrentListT[T]) RemoveRange(index, count int) {
	c.mux.Lock()
	defer c.mux.Unlock()

	newslice := make([]T, 0)
	newslice = append(newslice, c.data[0:index]...)
	newslice = append(newslice, c.data[index+count:len(c.data)]...)

	c.data = newslice
}
