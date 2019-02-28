package main

import "fmt"

var exp = "babad"

func main() {
	byteExp := []byte(exp)
	var longString string

	for i := 0; i < len(byteExp); i++ {
		mid := byteExp[i]
		tp := string(mid)
		for n := 1; (i + n) < len(byteExp)-1; n++ {
			if byteExp[i] == byteExp[i+n] {
				tp = tp + string(byteExp[i+n])
			}
		}
		if len(tp) > len(longString) {
			longString = tp
		}
		tp = string(mid)
		for n := 1; i-n >= 0; n++ {
			if (i + n) <= len(byteExp)-1 {
				if byteExp[i-n] == byteExp[i+n] {
					tp = string(byteExp[i-n]) + tp + string(byteExp[i+n])
				}
			}
		}
		if len(tp) > len(longString) {
			longString = tp
		}
	}
	fmt.Println(longString)
}
