package storage

// максимальная длина кэша
var maxCacheLen int = 100

// Структура КЭШа. тип - Связанный список
type Cache struct {
	Len  int
	Head *Item
	Tail *Item
}

type Item struct {
	Data float32
	Prev *Item
	Next *Item
}

// Добавляем метрику в кэш
func (c *Cache) Add(value float32) {
	node := &Item{Data: value, Prev: c.Head}

	if c.Len > 0 {
		c.Head.Next = node
	}

	c.Head = node

	if c.Len == 0 {
		c.Tail = c.Head
	}

	// Если превышен размер кэша (MaxLen), обновляем Tail - последний(нижний) элемент в списке
	if c.Len >= maxCacheLen {
		c.Tail.Prev = nil
		c.Tail = c.Tail.Next
	} else {
		c.Len++
	}
}
