package queueusingstacks

type MyQueue struct {
	Input  []int
	Output []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (mq *MyQueue) Push(x int) {
	mq.Input = append(mq.Input, x)
}

func (mq *MyQueue) Pop() int {
	mq.Peek()
	last := mq.Output[len(mq.Output)-1]
	mq.Output = mq.Output[:len(mq.Output)-1]
	return last
}

func (mq *MyQueue) Peek() int {
	if len(mq.Output) == 0 {
		for len(mq.Input) != 0 {
			last := mq.Input[len(mq.Input)-1]
			mq.Input = mq.Input[:len(mq.Input)-1]
			mq.Output = append(mq.Output, last)
		}
	}
	return mq.Output[len(mq.Output)-1]
}

func (mq *MyQueue) Empty() bool {
	return len(mq.Input) == 0 && len(mq.Output) == 0
}
