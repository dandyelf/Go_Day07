package coinsort

import (
	"slices"
)

func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// Fixed
func MinCoins2(val int, coins []int) []int {
	res := make([]int, 0)

	if !slices.IsSorted(coins) {
		slices.Sort(coins)
	}

	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

// Optimized
func MinCoins2Optimized(value int, coins []int) []int {
	result := []int{}
	if !slices.IsSorted(coins) {
		slices.Sort(coins)
	}
	coins = slices.Compact(coins)
	for i := len(coins) - 1; i >= 0 && value > 0; i-- {
		for value >= coins[i] {
			value -= coins[i]
			result = append(result, coins[i])
		}
	}
	return result
}
