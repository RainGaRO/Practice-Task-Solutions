package main

import "testing"

func TestLRUCache(t *testing.T) {
	cache := Constructor(1000)

	cache.Set(0, "test")

	s, _ := cache.Get(0)
	if s != "test" {
		t.Errorf("Ожидаем 1 - получили %s\n", s)
	}
}
