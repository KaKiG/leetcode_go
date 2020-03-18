package main

func main() {
	var stack MyStack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Push(4)
	stack.Pop()
	stack.Pop()
	stack.Empty()
}

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	var stack MyStack
	stack.queue1 = []int{}
	stack.queue2 = []int{}
	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	m := len(this.queue1)
	for i := 0; i < m-1; i++ {
		tmp := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, tmp)
	}
	res := this.queue1[0]
	this.queue1 = this.queue2
	this.queue2 = []int{}
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	m := len(this.queue1)
	for i := 0; i < m-1; i++ {
		tmp := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, tmp)
	}
	res := this.queue1[0]
	this.queue2 = append(this.queue2, res)
	this.queue1 = this.queue2
	this.queue2 = []int{}
	return res
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return (len(this.queue1) == 0)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
