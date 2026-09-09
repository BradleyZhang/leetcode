// Source : https://leetcode.com/problems/min-stack
// Author : BradleyZhang
// Date   : 2026-09-09

/*****************************************************************************************************
 *
 * Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
 *
 * Implement the MinStack class:
 *
 * 	MinStack() initializes the stack object.
 * 	void push(int value) pushes the element value onto the stack.
 * 	void pop() removes the element on the top of the stack.
 * 	int top() gets the top element of the stack.
 * 	int getMin() retrieves the minimum element in the stack.
 *
 * You must implement a solution with O(1) time complexity for each function.
 *
 * Example 1:
 *
 * Input
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * Output
 * [null,null,null,null,-3,null,0,-2]
 *
 * Explanation
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin(); // return -3
 * minStack.pop();
 * minStack.top();    // return 0
 * minStack.getMin(); // return -2
 *
 * Constraints:
 *
 * 	-2^31 <= val <= 2^31 - 1
 * 	Methods pop, top and getMin operations will always be called on non-empty stacks.
 * 	At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
 ******************************************************************************************************/
package minstack

type MinStack struct {
	stack []int
	mins  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= value {
		this.mins = append(this.mins, value)
	}
}

func (this *MinStack) Pop() {
	if this.mins[len(this.mins)-1] == this.stack[len(this.stack)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
