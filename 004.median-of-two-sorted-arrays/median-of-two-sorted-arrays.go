package main

import (
	"fmt"
	"sort"
)

var numArr1 = []int{1, 2}
var numArr2 = []int{3, 4}

func main() {
	j := append(numArr1, numArr2...)
	sort.Ints(j)
	jLen := len(j)
	var result interface{}
	midIndex := jLen / 2
	if jLen%2 == 0 {
		result = (float64(j[midIndex]) + float64(j[midIndex-1])) / 2
	} else {
		result = j[midIndex]
	}
	fmt.Println(result)
}
