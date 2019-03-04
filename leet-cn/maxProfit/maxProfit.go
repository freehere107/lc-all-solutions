package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
}

func maxProfit(prices []int) int {
	sum := 0
	buy := 0
	for i := 0; i < len(prices); i++ {
		if buy == 0 {
			buy = prices[i]
		} else {
			if prices[i] > buy {

			}
		}
	}
	return sum
}
