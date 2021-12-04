package cache

import (
	"container/list"
	"sync"
)

type Cache struct {
	list       *list.List
	entry      map[string]*list.Element
	mu         *sync.RWMutex
	maxEntries uint
}

func NewCache(maxEntries uint) *Cache {
	return &Cache{
		maxEntries: maxEntries,
		list:       list.New(),
		entry:      make(map[string]*list.Element),
		mu:         &sync.RWMutex{},
	}
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if element, ok := c.entry[key]; ok {
		return element.Value
	}

	return nil
}

func (c *Cache) Set(key string, val interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, ok := c.entry[key]; ok {
		element.Value = val
		c.list.MoveToFront(element)
		c.entry[key] = element
	} else {
		c.entry[key] = c.list.PushFront(val)
	}

	c.evict()

	return true
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.entry[key]; !ok {
		return false
	}

	c.list.Remove(c.entry[key])
	delete(c.entry, key)

	return true
}

func (c *Cache) Clear() {
	c.list.Init()
	c.entry = make(map[string]*list.Element)
	c.maxEntries = 0
}

func (c *Cache) Len() int {
	return c.list.Len()
}

func (c *Cache) isFull() bool {
	if c.maxEntries >= uint(c.Len()) {
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

	c.maxEntries -= 1

	return true
}
