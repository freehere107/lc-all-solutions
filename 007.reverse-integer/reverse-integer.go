package main

import (
	"strconv"
	"fmt"
)

func main() {
	reverse(-345)
}

func reverse(x int) {
	neg := false
	if x < 0 {
		x = 0 - x
		neg = true
	}
	strX := strconv.Itoa(x)
	var newX string
	for i, _ := range strX {
		newX = newX + string(strX[len(strX)-1-i])
	}
	fmt.Println(newX)
	newXInt, _ := strconv.Atoi(newX)
	if neg {
		newXInt = -newXInt
	}
	fmt.Println(newXInt)
}
