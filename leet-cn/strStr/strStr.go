package main

import "fmt"

func main() {
	haystack := "mississippi"
	needle := "pi"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == needle {
		return 0
	}
	needleLen := len(needle)
	for i := 0; i < len(haystack); i++ {
		if i+needleLen <= len(haystack) {
			fmt.Println(i,string(haystack[i:i+needleLen]))
			if string(haystack[i:i+needleLen]) == needle {
				return i
			}
		}
	}
	return -1
}
