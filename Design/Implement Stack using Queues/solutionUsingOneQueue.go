type MyStack struct {
	q1 []int
}

func Constructor() MyStack {
	q1 := make([]int, 0)

	return MyStack{
		q1: q1,
	}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)

	for i := 0; i < len(this.q1)-1; i++ {
		this.q1 = append(this.q1, this.q1[0])
		this.q1 = this.q1[1:]
	}
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