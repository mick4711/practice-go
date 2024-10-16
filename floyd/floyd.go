package floyd

import "fmt"

func Triangle(rowCount int) [][]int {
	for i := range rowCount {
		outer := [][]int{}
		inner := []int{}
		for j := range i {
			fmt.Println(i, j)
			inner = []int{1}
		}
		append(outer, inner)
	}
	return outer
}
