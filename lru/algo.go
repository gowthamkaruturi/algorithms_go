package lru

type evictionAlgo interface {
	evict(c *Cache) interface{}
	get(node *node, c *Cache)
	set(node *node, c *Cache)
	set_overwrite(node *node, value interface{}, c *Cache)
}

func createEvictioAlgo(algoType string) evictionAlgo {
	if algoType == "fifo" {
		return &fifo{}
	} else if algoType == "lru" {
		return &lru{}
	}
	return nil
}
