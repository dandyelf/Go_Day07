package coinsort_test

import (
	"leftrana/moneybag/ex00/coinsort"
	"testing"
)

var amout = 99999
var coinDenominations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
var want = 1011

// var table = []struct {
// 	amout             int
// 	coinDenominations []int
// 	want              []int
// }{
// 	{amout: 100},
// 	{amout: 1000},
// 	{amout: 74382},
// 	{amout: 382399},
// }

func sumCoins(coins []int) int {
	sum := 0
	for _, number := range coins {
		sum += number
	}
	return sum
}

// Easy test
func TestMinCoins(t *testing.T) {
	got := coinsort.MinCoins(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

func TestMinCoins2(t *testing.T) {
	got := coinsort.MinCoins2(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

func TestMinCoins2Optimized(t *testing.T) {
	got := coinsort.MinCoins2Optimized(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

// Duplicated test
func TestMinCoinsDuplicated(t *testing.T) {
	amout = 99999
	coinDenominations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 12, 13, 13, 13, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 99, 99}
	got := coinsort.MinCoins(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

func TestMinCoins2Duplicated(t *testing.T) {
	amout = 99999
	coinDenominations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 12, 13, 13, 13, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 99, 99}
	got := coinsort.MinCoins2(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

func TestMinCoins2OptimizedDuplicated(t *testing.T) {
	amout = 99999
	coinDenominations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 12, 13, 13, 13, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 99, 99}
	got := coinsort.MinCoins2Optimized(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

// Zero test
// Original func endless looping
func TestMinCoins2Zero(t *testing.T) {
	amout = 99999
	coinDenominations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 12, 13, 13, 13, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 99, 99}
	got := coinsort.MinCoins2(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

func TestMinCoins2OptimizedZero(t *testing.T) {
	amout = 99999
	coinDenominations = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 11, 12, 13, 13, 13, 13, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 99, 99}
	got := coinsort.MinCoins2Optimized(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

// Fail test
func TestMinCoinsFail(t *testing.T) {
	amout = 6
	coinDenominations = []int{1, 3, 4}
	// right: []int{3, 3}, wrong: []int{4, 1, 1}
	want = 2
	got := coinsort.MinCoins(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}

func TestMinCoins2Fail(t *testing.T) {
	amout = 6
	coinDenominations = []int{1, 3, 4}
	// right: []int{3, 3}, wrong: []int{4, 1, 1}
	want = 2
	got := coinsort.MinCoins2(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}
func TestMinCoins2OptimizedFail(t *testing.T) {
	amout = 6
	coinDenominations = []int{1, 3, 4}
	// right: []int{3, 3}, wrong: []int{4, 1, 1}
	want = 2
	got := coinsort.MinCoins2Optimized(amout, coinDenominations)
	if sumCoins(got) != amout || len(got) != want {
		t.Error("amout: ", amout, "coinDenominations", coinDenominations, "want: ", want, "got: ", len(got))
	}
}
