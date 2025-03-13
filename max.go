package main

import (
	"fmt"
	"sort"
)

func maxDifferance(nums []int) int {
	sort.Ints(nums)
	maxDiff := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[(i)-1]
		if diff > maxDiff {
			maxDiff=diff
		}
	}
	return maxDiff
}
func main() {
	nums := []int{3, 6, 9, 1}
	result := maxDifferance(nums)
	fmt.Println(result)
}
