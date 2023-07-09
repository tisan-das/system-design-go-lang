package behavioral_patterns

import "fmt"

type EvictionAlgo interface {
	Evict(*Cache)
}

type Cache struct {
	mappedElem  map[string]string
	algo        EvictionAlgo
	capacity    int
	maxCapacity int
}

func (cache *Cache) InitCache(algo EvictionAlgo) {
	cache.capacity = 0
	cache.maxCapacity = 2
	cache.algo = algo
	cache.mappedElem = make(map[string]string)
}

func (cache *Cache) Add(key, value string) {
	if cache.capacity == cache.maxCapacity {
		cache.algo.Evict(cache)
	}
	cache.capacity++
}

type FifoEvictionAlgo struct{}

func (algo *FifoEvictionAlgo) Evict(cache *Cache) {
	fmt.Println("Evicting through Fifo policy")
	cache.capacity--
}

type LSTEvictionAlgo struct{}

func (algo *LSTEvictionAlgo) Evict(cache *Cache) {
	fmt.Println("Evicting through LST policy")
	cache.capacity--
}
