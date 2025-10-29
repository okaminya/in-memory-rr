package main

import "fmt"

type Cache struct {
	data map[string]int
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]int),
	}
}

func (c *Cache) Add(name string, id int) {
	c.data[name] = id
}

func (c *Cache) Get(name string) (int, bool) {
	val, ok := c.data[name]
	return val, ok
}

func (c *Cache) Delete(name string) {
	delete(c.data, name)
}

func main() {
	cache := NewCache()

	cache.Add("okidoki", 44)

	value, ok := cache.Get("okidoki")

	fmt.Println(value, ok)

	cache.Delete("okidoki")

	value, ok = cache.Get("okidoki")
	fmt.Println(value, ok)
}
