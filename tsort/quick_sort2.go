package tsort

import (
	"fmt"
)

func QuickSort_b() {
	data := []int{30, 8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}
	quickSort_b3(data, 0, len(data)-1)
	fmt.Println(data)
}

func quickSort_b3(data []int, left, right int) {
	if left >= right {
		return
	}

	a, t := left, left
	z := right
	temp := data[t]

	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if z > t {
			data[t] = data[z]
			t = z
		}
		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp
	if left < t {
		quickSort_b3(data, left, t)
	}
	if t+1 < right {
		quickSort_b3(data, t+1, right)
	}

}

func quickSort_b1(data []int, left, right int) {
	if left >= right {
		return
	}

	s := left
	e := right
	t := left
	temp := data[t]

	for e > s {
		for data[e] >= temp && e > s {
			e--
		}
		if t < e {
			data[t] = data[e]
			data[e] = temp
			t = e
		}

		for data[s] <= temp && e > s {
			s++
		}
		if t > e {
			data[t] = data[s]
			data[s] = temp
			t = s
		}

	}

	if t > left {
		quickSort_b1(data, left, t)
	}
	if right > t+1 {
		quickSort_b1(data, t+1, right)
	}
}

func quickSort_b2(data []int, left, right int) {
	if left >= right {
		return
	}

	s := left
	t := s
	e := right

	temp := data[t]

	for s < e {
		for temp <= data[e] && s < e {
			e--
		}
		data[t] = data[e]
		data[e] = temp
		t = e

		for temp >= data[s] && s < e {
			s++
		}
		data[t] = data[s]
		data[s] = temp
		t = s
	}
	quickSort_b2(data, left, t)
	quickSort_b2(data, t+1, right)
}
