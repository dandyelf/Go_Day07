package coinsort_test

import (
	"leftrana/moneybag/ex00/coinsort"
	"reflect"
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
