package concurrentlist

import "sync"

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
