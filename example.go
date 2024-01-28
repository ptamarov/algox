package main

import (
	"algox/algox"
	"fmt"
)

func main() {
	w := algox.New(7, [][]int{
		{0, 3, 6},
		{0, 3},
		{3, 4, 6},
		{2, 4, 5},
		{1, 2, 5, 6},
		{1, 6},
	},
	)

	fmt.Println(w.Subsets())
	w.Solve() // two solutions found
}
