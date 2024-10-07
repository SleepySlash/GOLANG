package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	sports := make([][]int, n)
	for i := range sports {
		sports[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scanf("%d", &sports[i][j])
		}
	}
	fmt.Println(sports)

	var ans = Train(n-1, -1, sports)
	fmt.Println(ans)
}

func Train(n int, pre int, sports [][]int) int {
	if n == -1 {
		return 0
	}
	newPre := -1
	max := 0
	for ind := 0; ind < 3; ind++ {
		if ind != pre {
			if max < sports[n][ind] {
				max = sports[n][ind]
				newPre = ind
			}
		}
	}
	ans := max + Train(n-1, newPre, sports)
	return ans

}
