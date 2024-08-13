package coinsort

import (
	"math"
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
func MinCoins2Optimized(val int, coins []int) []int {
	dp := make([]int, val+1)
	usedCoins := make([]int, val+1)

	for i := 1; i <= val; i++ {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for _, denom := range coins {
		for j := denom; j <= val; j++ {
			if dp[j-denom]+1 < dp[j] {
				dp[j] = dp[j-denom] + 1
				usedCoins[j] = denom
			}
		}
	}

	if dp[val] == math.MaxInt32 {
		return []int{}
	}

	result := []int{}
	for val > 0 {
		result = append(result, usedCoins[val])
		val -= usedCoins[val]
	}

	return result
}
