package min_stack

type MinStack struct {
	stack 		[]int
	minStack 	[]int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.minStack) > 0 {
		min := this.minStack[len(this.minStack)-1]
		if  min > x {
			this.minStack = append(this.minStack, x)
		}else{
			this.minStack = append(this.minStack, min)
		}
	}else{
		this.minStack = append(this.minStack, x)
	}
}


func (this *MinStack) Pop()  {
	if len(this.stack) <= 0 {
		return
	}
	this.stack = this.stack[0:len(this.stack)-1]
	this.minStack = this.minStack[0:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	if len(this.stack) <= 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.minStack) <= 0 {
		return 0
	}
	return this.minStack[len(this.minStack)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
