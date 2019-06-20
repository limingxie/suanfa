package tqueue

import "fmt"

//堆栈
func Stackfunc() {
	var aa string = "hahahahahah"
	var bb string = "hahahaha"
	Stackfunc1(aa)
	Stackfunc1(bb)
}

func Stackfunc1(a string) {
	str := make(map[int]byte)
	for i := range a {
		str[i] = a[i]
	}

	strLen := len(str)
	mid := strLen / 2

	s := make(map[int]byte)
	for i := 0; i <= mid-1; i++ {
		s[i] = str[i]
	}
	top := len(s)

	var next int
	if strLen%2 == 0 {
		next = mid
	} else {
		next = mid + 1
	}

	for i := next; i < strLen; i++ {
		if s[top-1] != str[i] {
			break
		}
		top--
	}

	if top == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
