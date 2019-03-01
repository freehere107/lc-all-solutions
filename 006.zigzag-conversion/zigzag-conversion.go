package main

import (
	"fmt"
	"strings"
)

var ts = "PAYPALISHIRING"

func main() {
	fmt.Println(convert(ts, 3))
}

//P   A   H   N
//A P L S I I G
//Y   I   R

func convert(s string, nRows int) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	var ans []string
	step := 2*nRows - 2
	for i := 0; i < nRows; i++ {
		one := i
		two := -i
	T:
		for {
			if one < n || two < n {
				if 0 <= two && two < n && one != two && i != nRows-1 {
					ans = append(ans, string(s[two]))
				}
				if one < n {
					ans = append(ans, string(s[one]))
				}
				one += step
				two += step
			} else {
				break T
			}
		}

	}
	return strings.Join(ans, "")

}
