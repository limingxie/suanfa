package tsort

import "fmt"

func InsertionSort() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for i := 0; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[j] > data[i] {
				v := data[i]
				for k := i; k > j; k-- {
					data[k] = data[k-1]
				}
				data[j] = v
			}
		}
	}

	fmt.Println(data)
}
