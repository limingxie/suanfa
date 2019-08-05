package tsearch

import "fmt"

var n int = 5
var aa map[int]int
var book map[int]int

//Depth First Search, DFS
func DFS() {
	aa = make(map[int]int)
	for i := 1; i < 10; i++ {
		aa[i] = 0
	}

	book = make(map[int]int)
	for i := 1; i < 10; i++ {
		book[i] = 0
	}

	dfs(1)

}

func dfs(step int) {
	if step == n+1 {
		for j := 1; j <= n; j++ {
			fmt.Print(aa[j])
		}
		fmt.Println("")
		return
	}

	for i := 1; i <= n; i++ {
		if book[i] == 0 {
			aa[step] = i
			book[i] = 1
			dfs(step + 1)
			book[i] = 0
		}
	}
}
