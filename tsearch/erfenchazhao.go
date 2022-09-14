package tsearch

import "fmt"

var arr = []int{1, 12, 15, 29, 33, 37, 116, 222, 399, 422, 543, 1000}

func ErfenchazhaoTest(aa int) {
	result := Erfenchazhao(0, len(arr)-1, aa)
	fmt.Println(result)
}

func Erfenchazhao(start, end, aa int) int {
	fmt.Println(start, "|", end)
	c := (end + start) / 2
	if aa == arr[c] {
		return c
	}

	if start == end || end-1 == start {
		if aa == arr[end] {
			return end
		} else {
			return -1
		}
	}

	if aa > arr[c] {
		return Erfenchazhao(c, end, aa)
	} else {
		return Erfenchazhao(start, c, aa)
	}
}
