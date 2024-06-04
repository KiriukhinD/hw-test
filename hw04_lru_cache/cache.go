package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    list
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    list{},
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		item.Value = value
		c.queue.MoveToFront(item)
		return true
	}

	if c.queue.Len() == c.capacity {
		delete(c.items, c.queue.Back().Value.(Key))
		c.queue.Remove(c.queue.Back())
	}

	newItem := c.queue.PushFront(key)
	newItem.Value = value
	c.items[key] = newItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = list{}
	c.items = make(map[Key]*ListItem)
}
