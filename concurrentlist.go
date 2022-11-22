package concurrentlist

import (
	"errors"
	"fmt"
	"sync"
)

type ConcurrentList struct {
	data []interface{}
	mux  sync.Mutex
}

func New() *ConcurrentList {
	return &ConcurrentList{
		data: make([]interface{}, 0),
	}
}

func (c *ConcurrentList) Add(v interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = append(c.data, v)
}

func (c *ConcurrentList) AddRange(v []interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = append(c.data, v...)
}

func (c *ConcurrentList) Clear() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data = make([]interface{}, 0)
}

func (c *ConcurrentList) Remove(v interface{}) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	newslice := make([]interface{}, 0)
	isexist := false

	for _, d := range c.data {
		if d == v && !isexist {
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

func (c *ConcurrentList) RemoveRange(index, count int) {
	c.mux.Lock()
	defer c.mux.Unlock()

	newslice := make([]interface{}, 0)
	newslice = append(newslice, c.data[0:index]...)
	newslice = append(newslice, c.data[index+count:len(c.data)]...)

	c.data = newslice
}

// just get ,not remove
func (c *ConcurrentList) Get(index int) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if len(c.data) <= index {
		return nil, errors.New(fmt.Sprintf("index: %d out of bound,max len: %d", index, len(c.data)))
	}
	return c.data[index], nil
}

// get all and not remove
func (c *ConcurrentList) GetAll() []interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.data
}

// get one item and remove
func (c *ConcurrentList) Take(index int) (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if d, err := c.Get(index); err != nil {
		return d, err
	} else {
		if c.Remove(d) {
			return d, nil
		} else {
			return nil, errors.New("remove fail")
		}
	}
}
func (c *ConcurrentList) TakeAll() []interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()

	r := c.GetAll()
	c.Clear()
	return r
}
