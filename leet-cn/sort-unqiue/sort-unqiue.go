package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 4, 5, 5, 6, 6, 6}
	removeDuplicates(nums)

}

func removeDuplicates(nums []int) int {
	j := 0
	c := nums
	for i := 0; i < len(c); i++ {
		if c[i] != c[j] {
			j += 1
			fmt.Println(j)
			nums[j] = c[i]
		}
	}
	return j + 1
}
