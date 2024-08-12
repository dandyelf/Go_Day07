package coinsort_test

import (
	"leftrana/moneybag/ex00/coinsort"
	"strconv"
	"testing"
)

const u = "coin err"

func TestCreateImage(t *testing.T) {
	mass := coinsort.MinCoins(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12})
	if len(mass) == 0 {
		t.Error(u)
	}
	test := 0
	for _, v := range mass {
		test = test + v
	}
	if test != 10 {
		t.Error(u + strconv.Itoa(test))
	}
}

func TestMinCoins3(t *testing.T) {
	want := 1563412123342
	coinDenominations := []int{1000, 20, 200, 100, 2, 5000, 10, 5, 1, 50}
	res := coinsort.MinCoins3(want, coinDenominations)

	got := 0
	for _, v := range res {
		got += v
	}

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
