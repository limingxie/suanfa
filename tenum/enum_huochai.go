package tenum

import (
	"fmt"
	"strconv"
)

func Enum_huochai() {
	cc := make(map[int]int)

	cc[0] = 6
	cc[1] = 2
	cc[2] = 5
	cc[3] = 5
	cc[4] = 4
	cc[5] = 5
	cc[6] = 6
	cc[7] = 3
	cc[8] = 7
	cc[9] = 6

	var x, y int
	for x = 0; x < 1111; x++ {
		for y = 0; y < 1111; y++ {
			if Enum3(x)+Enum3(y)+Enum3(x+y) == 14 {
				fmt.Println("====================")
				fmt.Print(Enum3(x))
				fmt.Print("  ")
				fmt.Print(Enum3(y))
				fmt.Print("  ")
				fmt.Print(Enum3(x + y))
				fmt.Println("")
				rr := strconv.Itoa(x) + " + " + strconv.Itoa(y) + " = " + strconv.Itoa(x+y)
				fmt.Println(rr)
			}
		}
	}

}

func Enum3(x int) int {

	cc := make(map[int]int)

	cc[0] = 6
	cc[1] = 2
	cc[2] = 5
	cc[3] = 5
	cc[4] = 4
	cc[5] = 5
	cc[6] = 6
	cc[7] = 3
	cc[8] = 7
	cc[9] = 6

	var count int
	for x/10 != 0 {
		count = count + cc[x%10]
		x = x / 10
	}

	count = count + cc[x]

	return count
}
