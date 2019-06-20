package tenum

import (
	"fmt"
	"strconv"
)

//枚举
func Enum1() {

	var total int

	book := make(map[int]int)
	for i := 1; i < 10; i++ {
		book[i] = 0
	}

	aa := make(map[int]int)
	for i := 1; i < 10; i++ {
		aa[i] = 0
	}

	for aa[1] = 1; aa[1] < 10; aa[1]++ {
		for aa[2] = 1; aa[2] < 10; aa[2]++ {
			for aa[3] = 1; aa[3] < 10; aa[3]++ {
				for aa[4] = 1; aa[4] < 10; aa[4]++ {
					for aa[5] = 1; aa[5] < 10; aa[5]++ {
						for aa[6] = 1; aa[6] < 10; aa[6]++ {
							for aa[7] = 1; aa[7] < 10; aa[7]++ {
								for aa[8] = 1; aa[8] < 10; aa[8]++ {
									for aa[9] = 1; aa[9] < 10; aa[9]++ {

										for i := 1; i < 10; i++ {
											// fmt.Println("===========aa[" + strconv.Itoa(i) + "]============")
											// fmt.Println(aa[i])
											book[aa[i]] = book[aa[i]] + 1
										}

										var temp int = 0
										for i := range book {
											if book[i] == 1 {
												temp++
											}
										}
										if temp == 9 {
											if aa[1]*100+aa[2]*10+aa[3]+aa[4]*100+aa[5]*10+aa[6] == aa[7]*100+aa[8]*10+aa[9] {
												total++
												// fmt.Println("=======================")

												// rr := strconv.Itoa(aa[1]) + strconv.Itoa(aa[2]) + strconv.Itoa(aa[3]) + " + " +
												// 	strconv.Itoa(aa[4]) + strconv.Itoa(aa[5]) + strconv.Itoa(aa[6]) + " = " +
												// 	strconv.Itoa(aa[7]) + strconv.Itoa(aa[8]) + strconv.Itoa(aa[9])
												// fmt.Println(rr)
											}
										}

										for i := 1; i < 10; i++ {
											book[i] = 0
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(total / 2)
}

func AAA() {

	book := make(map[int]int)
	for i := 1; i < 10; i++ {
		book[i] = 1
	}

	// fmt.Println(book)
	for j := 1; j < 3; j++ {
		fmt.Println("===========j[" + strconv.Itoa(j) + "]============")
		for i := range book {
			if book[i] != 1 {
				break
			} else {
				fmt.Println("===========book[" + strconv.Itoa(i) + "]============")
				fmt.Println(book[i])

			}
		}
	}
}
