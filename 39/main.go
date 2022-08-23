package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8)) //nolint:gomnd
}

func combinationSum(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)

	rec(target, candidates, 0, nil, &res)

	return
}

func rec(left int, candidates []int, pos int, tmp []int, res *[][]int) {
	if left == 0 {
		*res = append(*res, append([]int{}, tmp...))

		return
	}

	for i := pos; i < len(candidates); i++ {
		if left-candidates[i] < 0 {
			return
		}

		rec(left-candidates[i], candidates, i, append(tmp, candidates[i]), res)
	}
}
