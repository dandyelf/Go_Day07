package coinsort_test

import (
	"leftrana/moneybag/ex00/coinsort"
	"strconv"
	"testing"
)

const u = "coin err"

func TestMinCoins(t *testing.T) {
	mass := coinsort.MinCoins(15, []int{1, 5, 10, 50, 100, 500, 1000})
	if len(mass) == 0 {
		t.Error(u)
	}
	test := 0
	for _, v := range mass {
		test = test + v
	}
	if test != 15 {
		t.Error(u + strconv.Itoa(test))
	}
	mass = coinsort.MinCoins(15, []int{1, 3, 4, 7, 13, 15})
	if len(mass) == 0 {
		t.Error(u)
	}
	test = 0
	for _, v := range mass {
		test = test + v
	}
	if test != 15 {
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
