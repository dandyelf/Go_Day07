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
	coins = slices.Compact(coins)

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
/*
func MinCoins2Optimized(val int, coins []int) []int {
	res := make([]int, 0)

	if !slices.IsSorted(coins) {
		slices.Sort(coins)
	}
	coins = slices.Compact(coins)

	for i := len(coins) - 1; i >= 0; i-- {

		if val >= coins[i] {
			qty := val / coins[i]
			val = val % coins[i]

			for k := 1; k <= qty; k++ {
				res = append(res, coins[i])
			}
		}
	}

	return res
}
*/

func MinCoins2Optimized(val int, coins []int) []int {
	res := make([]int, 0)
	for i := len(coins) - 1; i >= 0; i-- {
		count := val / coins[i]
		if count > 0 {
			// Добавляем нужное количество монет в результат
			for j := 0; j < count; j++ {
				res = append(res, coins[i])
			}
			val -= count * coins[i]
		}
	}
	return res
}
