package tqueue

import "fmt"

//队列
type queue struct {
	data map[int]int
	head int
	tail int
}

func Queuefunc() {
	var q queue

	q.head = 0
	q.tail = 9
	q.data = make(map[int]int)
	q.data[0] = 6
	q.data[1] = 3
	q.data[2] = 1
	q.data[3] = 7
	q.data[4] = 5
	q.data[5] = 8
	q.data[6] = 9
	q.data[7] = 2
	q.data[8] = 4

	for i := 0; i < len(q.data); i++ {
		// fmt.Print(strconv.Itoa(i) + ":")
		fmt.Print(q.data[i])
		fmt.Print("  ")
	}
	fmt.Println("")
	fmt.Println("---------------------------------------")

	for q.head < q.tail {
		fmt.Print(q.data[q.head])
		fmt.Print("  ")
		q.head++
		q.data[q.tail] = q.data[q.head]
		q.tail++
		q.head++
	}

	fmt.Println("")
}
