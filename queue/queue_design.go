package main

func testQueue() {
	//queue := Queue{}
	//queue.Push(1)
	//queue.Push(2)
	//fmt.Println(queue.Pop())
	//fmt.Println(queue.isEmpty())

	//a := []int{1, 2}
	//b := []int{3, 4}
	//c := append(a, b...)
	//fmt.Println(c)
}

type Queue []int

func (queue *Queue) Push(value int) {
	*queue = append(*queue, value)
}

func (queue *Queue) Pop() int {
	popValue := (*queue)[0]
	*queue = (*queue)[1:]
	return popValue
}

func (queue Queue) isEmpty() bool {
	return len(queue) == 0
}
