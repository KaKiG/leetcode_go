package main

func main() {

}

type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	var queue MyQueue
	queue.stack1 = []int{}
	queue.stack2 = []int{}
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	m := len(this.stack1)
	for i := 0; i < m-1; i++ {
		tmp := this.stack1[len(this.stack1)-1]
		this.stack1 = this.stack1[:len(this.stack1)-1]
		this.stack2 = append(this.stack2, tmp)
	}
	res := this.stack1[0]
	this.stack1 = []int{}
	for i := 0; i < m-1; i++ {
		tmp := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		this.stack1 = append(this.stack1, tmp)
	}
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	m := len(this.stack1)
	for i := 0; i < m-1; i++ {
		tmp := this.stack1[len(this.stack1)-1]
		this.stack1 = this.stack1[:len(this.stack1)-1]
		this.stack2 = append(this.stack2, tmp)
	}
	res := this.stack1[0]
	for i := 0; i < m-1; i++ {
		tmp := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		this.stack1 = append(this.stack1, tmp)
	}
	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return (len(this.stack1) == 0)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
