package problem0121

func maxProfit_1(prices []int) int {
	maxprofit := 0;

	for i := 0; i < len(prices) - 1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i];
			if (profit > maxprofit) {
				maxprofit = profit;
			}
		}
	}

	return maxprofit;
}

func maxProfit_2(prices []int) int {
	minprice := int(^uint(0) >> 1)

	maxprofit := 0;

	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i];
		} else if prices[i] - minprice > maxprofit {
			maxprofit = prices[i] - minprice;
		}
	}

	return maxprofit;
}