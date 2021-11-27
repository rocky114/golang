package cache

import "container/list"

type Cache struct {
	maxBytes int
	list     *list.List
	entry    map[string]*list.Element
}

func New(size int) *Cache {
	return &Cache{
		maxBytes: size,
		list:     list.New(),
		entry:    make(map[string]*list.Element),
	}
}

func (c *Cache) Get(key string) interface{} {
	if element, ok := c.entry[key]; ok {
		return element.Value
	}

	return nil
}

func (c *Cache) Set(key string, val interface{}) bool {
	c.evict()

	element := c.list.PushFront(val)
	c.entry[key] = element

	return true
}

func (c *Cache) isFull() bool {
	if c.list.Len() >= c.maxBytes {
		return true
	}

	return false
}

func (c *Cache) evict() bool {
	if c.isFull() == false {
		return true
	}

	c.list.Remove(c.list.Back())
	c.purge()

	return true
}

func (c *Cache) purge() {
	for key, val := range c.entry {
		if val == nil {
			delete(c.entry, key)
		}
	}
}
