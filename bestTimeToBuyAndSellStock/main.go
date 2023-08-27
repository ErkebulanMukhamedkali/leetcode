package main

import "log"

func main() {
	log.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
    profit := 0
    buy, sell := 0, 1
    for sell <= len(prices)-1 {
        if prices[sell] < prices[buy] {
            buy = sell
            sell ++
            continue
        }
        local_profit := prices[sell]-prices[buy]
        if local_profit > profit {
            profit = local_profit
        }
        sell++
    }
    return profit
}
