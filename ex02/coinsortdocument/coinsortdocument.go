// This package provides functions for finding the minimum number of coins needed to make a given value.
/*
For get this doccument, do: 'godoc -url "/pkg/leftrana/moneybag/ex02/coinsortdocument/" > doc.html'
First, make sure you have GOPATH=$HOME/go and $GOPATH/bin in your PATH.
Next 'go install golang.org/x/tools/cmd/godoc' to install the godoc tool.
*/
package coinsortdocument

import (
	"math"
	"slices"
)

// Original
/*
This function finds the minimum number of coins needed to make a given value.
It does this by iterating over the coins in reverse order and subtracting each coin from the value as many times as possible.
The resulting coins are appended to a result slice, which is then returned.
*/
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
/*
This function finds the minimum number of coins needed to make a given value using dynamic programming. It first initializes a dp slice and a usedCoins slice.
The dp slice stores the minimum number of coins needed to make each value from 0 to val.
The usedCoins slice stores the coin used to make each value.
The function then iterates over each coin and for each coin, it iterates over each value from the coin value to val.
If the number of coins needed to make the current value using the current coin is less than the current number of coins needed to make the current value,
it updates the dp slice and the usedCoins slice.

Finally, if the dp slice for the final value is still the maximum integer, it returns an empty slice.
Otherwise, it iterates over the usedCoins slice to build the result slice and returns it.
*/
func MinCoins2(val int, coins []int) []int {
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

// Optimized
/*
This function is an optimized version of MinCoins2.
It first checks if the coins slice is sorted and if not, it sorts it. It then removes any duplicate coins from the slice.

The function then proceeds as MinCoins2, but with the optimized coins slice.
*/
func MinCoins2Optimized(val int, coins []int) []int {
	if !slices.IsSorted(coins) {
		slices.Sort(coins)
	}
	coins = slices.Compact(coins)
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
