package coinsort

import (
	"slices"
	"sort"
	"sync"
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

// Optimized with goroutines ------------------------------------------------------

// Result struct
type Result struct {
	mu    sync.Mutex
	wg    sync.WaitGroup
	slice []int
}

// Add the method to appent coin
func (r *Result) Add(coin int, qty int) {
	r.mu.Lock()
	for k := 1; k <= qty; k++ {
		r.slice = append(r.slice, coin)
		r.wg.Done()
	}
	r.mu.Unlock()
}

// Function
func MinCoins4(val int, coins []int) []int {
	res := Result{slice: make([]int, 0)}

	if !slices.IsSorted(coins) {
		slices.Sort(coins)
	}
	coins = slices.Compact(coins)

	for i := len(coins) - 1; i >= 0; i-- {

		if val >= coins[i] {
			qty := val / coins[i]
			val = val % coins[i]

			res.wg.Add(qty)
			go func() {
				res.Add(coins[i], qty)
			}()
		}
	}

	res.wg.Wait()
	sort.Sort(sort.Reverse(sort.IntSlice(res.slice)))
	return res.slice
}
