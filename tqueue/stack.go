package tqueue

import "fmt"

func Stackfunc() {
	var aa string = "hahahahah"
	// var bb string = "hahahaha"
	Stackfunc1(aa)
	// Stackfunc1(bb)
}

func Stackfunc1(a string) {
	str := make(map[int]byte)
	for i := range a {
		str[i] = a[i]
	}
	fmt.Println(str)

	leng := len(str)
	mid := leng/2 - 1

	s := make(map[int]byte)
	for i := 0; i <= mid; i++ {
		s[i] = str[i]
	}
	top := len(s)
	fmt.Println(s)

	// for i := 0; i < len(s); i++ {
	// 	fmt.Print(s[i])+ 1
	// 	fmt.Print("  ")
	// }
	// fmt.Println("")
	// fmt.Println("---------------------------------------")

	var next int
	if leng%2 == 0 {
		next = mid
	} else {
		next = mid + 1
	}
	fmt.Println(next)

	for i := next; i < mid; i++ {
		if s[top] != str[i] {
			break
		}
		top--
	}
	fmt.Println(top)
	if top == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
