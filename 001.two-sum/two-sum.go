package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(solve(nums, target))
}

func solve(nums []int, target int) []int {
	for index, value := range nums {
		otherNum := target - value
		if inArray(otherNum, nums) != -1 {
			return []int{index, inArray(otherNum, nums)}
		}
	}
	return []int{}
}

func inArray(noodle int, arr []int) int {
	for k, v := range arr {
		if noodle == v {
			return k
		}
	}
	return -1
}
