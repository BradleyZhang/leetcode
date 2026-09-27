// Source : https://leetcode.com/problems/lru-cache
// Author : BradleyZhang
// Date   : 2026-09-27

/*****************************************************************************************************
 *
 * Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
 *
 * Implement the LRUCache class:
 *
 * 	LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
 * 	int get(int key) Return the value of the key if the key exists, otherwise return -1.
 * 	void put(int key, int value) Update the value of the key if the key exists. Otherwise, add
 * the key-value pair to the cache. If the number of keys exceeds the capacity from this operation,
 * evict the least recently used key.
 *
 * The functions get and put must each run in O(1) average time complexity.
 *
 * Example 1:
 *
 * Input
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * Output
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * Explanation
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // cache is {1=1}
 * lRUCache.put(2, 2); // cache is {1=1, 2=2}
 * lRUCache.get(1);    // return 1
 * lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
 * lRUCache.get(2);    // returns -1 (not found)
 * lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
 * lRUCache.get(1);    // return -1 (not found)
 * lRUCache.get(3);    // return 3
 * lRUCache.get(4);    // return 4
 *
 * Constraints:
 *
 * 	1 <= capacity <= 3000
 * 	0 <= key <= 10^4
 * 	0 <= value <= 10^5
 * 	At most 2 * 10^5 calls will be made to get and put.
 ******************************************************************************************************/

package lrucache

type Node struct {
	prev *Node
	next *Node
	val  int
	key  int
}

type LRUCache struct {
	key2node map[int]*Node
	cap      int
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	tail := &Node{}
	head := &Node{next: tail}
	tail.prev = head
	return LRUCache{
		key2node: make(map[int]*Node, capacity),
		cap:      capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.key2node[key]; ok {
		this.up(n)

		return n.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.key2node[key]; ok {
		this.up(n)
		n.val = value
		return
	}
	if len(this.key2node) >= this.cap {
		this.pop()
	}
	node := &Node{
		val: value,
		key: key,
	}
	this.up(node)
	this.key2node[key] = node
}
func (this *LRUCache) up(n *Node) {
	if n.prev != nil && n.next != nil {
		n.prev.next = n.next
		n.next.prev = n.prev
	}

	this.head.next.prev = n
	n.next = this.head.next
	this.head.next = n
	n.prev = this.head
}
func (this *LRUCache) pop() {
	if this.tail.prev != this.head {
		delete(this.key2node, this.tail.prev.key)
		this.tail.prev.prev.next = this.tail
		this.tail.prev = this.tail.prev.prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
