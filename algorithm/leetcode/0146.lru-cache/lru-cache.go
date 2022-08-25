package problem0146

import (
	"container/list"
	"fmt"
)

// LRUCache: 基于HashMap + 双向链表实现
type LRUCache struct {
	cap int
	l   *list.List				// 双向链表, 链表头部存放访问时间最晚的数据节点，链表尾部存放访问时间最早的数据节点
	m   map[int]*list.Element
}

// 双向链表中的数据节点
type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	// 读缓存，若缓存命中, 则将命中元素移动到链表头部, 并返回该元素; 否则返回-1
	if node, ok := c.m[key]; ok {
		// node.Value.(*list.Element): 根据key从map中获取对应节点, 并强转为*list.Element类型
		// node.Value.(*list.Element).Value.(Pair): 获取*list.Element节点的Value, 并强转为Pair类型
		// node.Value.(*list.Element).Value.(Pair).value: 从Pair中获取value属性
		val := node.Value.(*list.Element).Value.(Pair).value
		c.l.MoveToFront(node)
		fmt.Println("读缓存，缓存命中，移动数据节点到链表头部，并返回数据节点:",  val)
		return val
	}

	fmt.Println("读缓存，缓存未命中，返回-1")
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		// 写缓存，若缓存命中, 则更新链表头部元素
		c.l.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
		fmt.Println("写缓存，缓存命中，更新链表头部元素")
	} else {
		// 写缓存，若缓存未命中且缓存已满, 则发生LRU替换, 删除最近最久未被访问元素(链表尾部)
		if c.l.Len() == c.cap {
			listBackEle := c.l.Back()
			idx := listBackEle.Value.(*list.Element).Value.(Pair).key
			delete(c.m, idx)				// hashMap删索引
			c.l.Remove(listBackEle)			// 链表删节点
			fmt.Println("写缓存，缓存未命中且缓存已满，删除链表尾部元素:", listBackEle.Value.(*list.Element).Value.(Pair).value)
		}
		// 写缓存，若缓存未命中且缓存未满，则直接插入新数据节点到链表头部
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		ptr := c.l.PushFront(node)			// 链表添节点
		c.m[key] = ptr						// hashMap添索引
		fmt.Println("写缓存，缓存未命中且缓存未满，插入数据节点到链表头部:", value)
	}
}
