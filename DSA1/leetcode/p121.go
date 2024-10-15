package leetcode

//0(n) time complexity, 0(1) space complexity
func MaxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}

//0(n) time complexity, 0(n) space complexity
func MaxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	profit := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > profit[i-1] {
			profit[i] = prices[i] - min
		} else {
			profit[i] = profit[i-1]
		}
	}

	return profit[len(prices)-1]
}