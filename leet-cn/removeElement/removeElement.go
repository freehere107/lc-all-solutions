package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val) == 2)
}
func removeElement(nums []int, val int) int {
	j := 0
	c := nums
	for i := 0; i < len(c); i++ {
		if c[i] != val {
			nums[j] = c[i]
			j += 1
		}
	}
	return j

}
