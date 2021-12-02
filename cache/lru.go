package cache

import (
	"container/list"
	"reflect"
	"sync"
	"unsafe"
)

type Cache struct {
	total uint
	used  uint
	list  *list.List
	entry map[string]*list.Element
	mu    *sync.Mutex
}

func NewCache(total uint) *Cache {
	return &Cache{
		total: total,
		list:  list.New(),
		entry: make(map[string]*list.Element),
		mu:    &sync.Mutex{},
	}
}

func (c *Cache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.entry[key]; ok {
		return element.Value
	}

	return nil
}

func (c *Cache) Set(key string, val interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.free(key)

	if element, ok := c.entry[key]; ok {
		element.Value = val
		c.list.MoveToFront(element)
		c.entry[key] = element
	} else {
		c.entry[key] = c.list.PushFront(val)
	}

	c.used += uint(len(key)) + length(val)

	c.evict()

	return true
}

func length(val interface{}) uint {
	kind := reflect.TypeOf(val).Kind()
	if kind == reflect.Array || kind == reflect.String || kind == reflect.Chan || kind == reflect.Map || kind == reflect.Slice {
		return uint(reflect.ValueOf(val).Len())
	}

	return uint(unsafe.Sizeof(val))
}

func (c *Cache) free(key string) {
	if _, ok := c.entry[key]; !ok {
		return
	}

	c.used -= length(key)

	c.used -= length(c.entry[key].Value)
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.free(key)

	if _, ok := c.entry[key]; !ok {
		return false
	}

	c.list.Remove(c.entry[key])
	delete(c.entry, key)

	return true
}

func (c *Cache) Len() int {
	return c.list.Len()
}

func (c *Cache) isFull() bool {
	if c.used >= c.total {
		return true
	}

	return false
}

func (c *Cache) evict() bool {
	if c.isFull() == false {
		return true
	}

	c.list.Remove(c.list.Back())
	for key, val := range c.entry {
		if val == nil {
			delete(c.entry, key)
		}
	}

	return true
}
