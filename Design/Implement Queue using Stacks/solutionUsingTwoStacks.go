type MyQueue struct {
	s1 []int
	s2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.s2 = append(this.s2, x)
}

func (this *MyQueue) Pop() int {
	if len(this.s1) == 0 {
		this.transfer()
	}

	topIndexElement := len(this.s1) - 1
	topElement := this.s1[topIndexElement]
	this.s1 = this.s1[:topIndexElement]

	return topElement
}

func (this *MyQueue) Peek() int {
	if len(this.s1) == 0 {
		this.transfer()
	}

	topIndexElement := len(this.s1) - 1
	return this.s1[topIndexElement]
}

func (this *MyQueue) Empty() bool {
	if len(this.s1) == 0 && len(this.s2) == 0 {
		return true
	}

	return false
}

func (this *MyQueue) transfer() {
	for len(this.s2) > 0 {
		topIndexElement := len(this.s2) - 1
		this.s1 = append(this.s1, this.s2[topIndexElement])
		this.s2 = this.s2[:topIndexElement]
	}
}