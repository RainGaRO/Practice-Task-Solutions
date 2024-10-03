package main

import (
	"container/list"
	"fmt"
)

// Хранение ключа и кэшируемого элемента
type KeyValue struct {
	Key   int
	Value string
}

// Кол-во ячеек, двусвязный список, хэш-таблица
type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRUCache) Get(key int) (string, bool) {
	if n, ok := c.items[key]; ok {
		c.queue.MoveToFront(n)
		return n.Value.(KeyValue).Value, true
	}
	return "", false
}

func (c *LRUCache) Set(key int, value string) {
	kv := KeyValue{key, value}

	if n, ok := c.items[key]; ok {
		n.Value = kv
		c.queue.MoveToFront(n)
		c.items[key] = n
	} else {
		c.items[key] = c.queue.PushFront(kv)
	}
	//если очередь больше макс. кол-ва ячеек удаляем последни элемент и ключ из хэш-таблицы
	if len(c.items) > c.capacity {
		delete(c.items, c.queue.Back().Value.(KeyValue).Key)
		c.queue.Remove(c.queue.Back())
	}

}

func main() {
	cache := Constructor(1000)

	cache.Set(0, "test")
	cache.Set(1, "test")
	cache.Set(2, "test")
	cache.Set(3, "test")
	cache.Set(4, "test")

	fmt.Println(cache)

	v, exists := cache.Get(0)
	if exists {
		fmt.Println(v)
	} else {
		fmt.Println("Элемент не найден")
	}
}
