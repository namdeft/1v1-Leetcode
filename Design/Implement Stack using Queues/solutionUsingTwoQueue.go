type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	q1 := make([]int, 0)
	q2 := make([]int, 0)

	return MyStack{
		q1: q1,
		q2: q2,
	}
}

func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)

	for len(this.q1) > 0 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}

	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStack) Pop() int {
	firstElement := this.q1[0]
	this.q1 = this.q1[1:]

	return firstElement
}

func (this *MyStack) Top() int {
	return this.q1[0]
}

func (this *MyStack) Empty() bool {
	if len(this.q1) == 0 {
		return true
	}

	return false
}