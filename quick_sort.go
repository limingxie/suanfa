package main

import "fmt"

var qq map[int]int

func QuickSort() {
	qq = make(map[int]int)
	qq[0] = 6
	qq[1] = 1
	qq[2] = 2
	qq[3] = 7
	qq[4] = 9
	qq[5] = 3
	qq[6] = 4
	qq[7] = 5
	qq[8] = 10
	qq[9] = 8

	QuickSort1(qq[1], qq[9])

	fmt.Println(qq)
}

func QuickSort1(left int, right int) {
	if left > right {
		return
	}

}
