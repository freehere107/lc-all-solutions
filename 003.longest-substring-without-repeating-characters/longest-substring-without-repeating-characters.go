package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "pwwkew"
	result := findRep(str)
	fmt.Print(result)
}

func findRep(str string) string {
	strArr := strings.Split(str, "")
	strMap := map[string]string{}
	oldArr := []string{}
	newArr := []string{}
	for _, v := range strArr {
		if strMap[v] == "1" {
			strMap = map[string]string{v: "1"}
			newArr = []string{v}
			if len(newArr) > len(oldArr) {
				oldArr = newArr
			}
		} else {
			strMap[v] = "1"
			newArr = append(newArr, v)
			if len(newArr) > len(oldArr) {
				oldArr = newArr
			}
		}

	}
	return strings.Join(oldArr, "")
}
