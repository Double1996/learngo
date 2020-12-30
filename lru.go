package lru

import "container/list"

type LruCache struct {
	Capacity 	int 
	cache 		map[int]*list.Element
	list 			*list.List
}

type Pair struct {
	key 	int
	value int 
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) int {
		if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if this.list.Len() >= this.capacity {
			delete(this.cache,this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	}
}