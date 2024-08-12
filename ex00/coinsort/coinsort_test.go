package coinsort_test

import (
	"leftrana/moneybag/ex00/coinsort"
	"reflect"
	"testing"
)

// const u = "coin err"

// func TestMinCoins1(t *testing.T) {
// 	mass := coinsort.MinCoins(15, []int{1, 5, 10, 50, 100, 500, 1000})
// 	if len(mass) == 0 {
// 		t.Error(u)
// 	}
// 	test := 0
// 	for _, v := range mass {
// 		test = test + v
// 	}
// 	if test != 15 {
// 		t.Error(u + strconv.Itoa(test))
// 	}
// 	mass = coinsort.MinCoins(15, []int{1, 3, 4, 7, 13, 15})
// 	if len(mass) == 0 {
// 		t.Error(u)
// 	}
// 	test = 0
// 	for _, v := range mass {
// 		test = test + v
// 	}
// 	if test != 15 {
// 		t.Error(u + strconv.Itoa(test))
// 	}
// }

var amout = 99999
var coinDenominations = []int{1000, 20, 200, 100, 2, 5000, 10, 5, 50, 1, 10000, 50000, 20000}
var want = []int{50000, 20000, 20000, 5000, 1000, 1000, 1000, 1000, 200, 200, 200, 200, 100, 50, 20, 20, 5, 2, 2}

func TestMinCoins(t *testing.T) {
	got := coinsort.MinCoins(amout, coinDenominations)

	if !reflect.DeepEqual(got, want) {
		t.FailNow()
	}
}

func TestMinCoins2(t *testing.T) {
	got := coinsort.MinCoins2(amout, coinDenominations)

	if !reflect.DeepEqual(got, want) {
		t.FailNow()
	}
}

func TestMinCoins3(t *testing.T) {
	got := coinsort.MinCoins3(amout, coinDenominations)

	if !reflect.DeepEqual(got, want) {
		t.FailNow()
	}
}
func TestMinCoins4(t *testing.T) {
	got := coinsort.MinCoins4(amout, coinDenominations)

	if !reflect.DeepEqual(got, want) {
		t.FailNow()
	}
}

func TestSpeed(t *testing.T) {
	want := 1563412123342
	coinDenominations := []int{1000, 20, 200, 100, 2, 5000, 10, 5, 1, 50}
	res := coinsort.MinCoins(want, coinDenominations)

	got := 0
	for _, v := range res {
		got += v
	}

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
