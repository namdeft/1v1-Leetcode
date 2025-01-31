type LRUCache struct {
	capacity    int
	mp          map[int]*LinkedListNode
	top, bottom *LinkedListNode
}

type LinkedListNode struct {
	key, value int
	next, prev *LinkedListNode
}

func Constructor(capacity int) LRUCache {
	top := &LinkedListNode{}
	bottom := &LinkedListNode{}

	top.prev = bottom
	bottom.next = top

	return LRUCache{
		capacity: capacity,
		mp:       make(map[int]*LinkedListNode),
		top:      top,
		bottom:   bottom,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.mp[key]; !ok {
		return -1
	} else {
		this.moveNodeToTop(v)

		return this.mp[key].value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.mp[key]; ok {
		v.value = value
		this.moveNodeToTop(v)
	} else {
		if len(this.mp) == this.capacity {
			bottom := this.bottom.next
			this.bottom.next = bottom.next
			this.bottom.next.prev = this.bottom
			delete(this.mp, bottom.key)
		}

		newNode := &LinkedListNode{
			key:   key,
			value: value,
		}
		this.addNodeToTop(newNode)
		this.mp[key] = newNode
	}
}

func (this *LRUCache) moveNodeToTop(node *LinkedListNode) {
	node.next.prev = node.prev
	node.prev.next = node.next

	this.addNodeToTop(node)
}

func (this *LRUCache) addNodeToTop(node *LinkedListNode) {
	node.next = this.top
	node.prev = this.top.prev
	this.top.prev.next = node
	this.top.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */