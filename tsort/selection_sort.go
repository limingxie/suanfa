package tsort

import "fmt"

//选择排序
func SelectionSort() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	var book int
	for i := 0; i < len(data); i++ {
		v := data[i]
		for j := i; j < len(data); j++ {
			if v >= data[j] {
				book = j
				v = data[j]
			}
		}
		data[book] = data[i]
		data[i] = v
	}

	fmt.Println(data)
}

//选择排序
func SelectionSort1() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	var book int
	for i := 0; i < len(data); i++ {
		book = i
		for j := i; j < len(data); j++ {
			if data[book] > data[j] {
				book = j
			}
		}
		v := data[i]
		data[i] = data[book]
		data[book] = v
	}

	fmt.Println(data)
}
