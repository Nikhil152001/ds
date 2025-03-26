package main

import "fmt"

func majorityElement(nums []int) int {
	candidate, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
		if count == 0 {
			candidate = nums[i]
			count = 0
		}
	}
	return candidate
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println("majorityElement:", majorityElement(nums))
}
