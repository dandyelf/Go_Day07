package coinsort_test

import (
	"leftrana/moneybag/ex00/coinsort"
	"reflect"
	"testing"
)

var amout = 99999
var coinDenominations = []int{1000, 20, 200, 100, 2, 5000, 10, 5, 50, 1, 10000, 50000, 20000}
var want = []int{50000, 20000, 20000, 5000, 1000, 1000, 1000, 1000, 200, 200, 200, 200, 100, 50, 20, 20, 5, 2, 2}

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

func TestMinCoins(t *testing.T) {
	got := coinsort.MinCoins(amout, coinDenominations)
	if !reflect.DeepEqual(got, want) {
		t.Error("first func")
	}
}

func TestMinCoins2(t *testing.T) {
	got := coinsort.MinCoins2(amout, coinDenominations)

	if !reflect.DeepEqual(got, want) {
		t.FailNow()
	}
}

func TestMinCoins2Optimized(t *testing.T) {
	got := coinsort.MinCoins2Optimized(amout, coinDenominations)

	if !reflect.DeepEqual(got, want) {
		t.FailNow()
	}
}
