package fib

import (
	"math/big"
)

var (
	zero = big.NewInt(0)
	one = big.NewInt(1)
	two = big.NewInt(2)
)

func Fib(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	if v, ok := memo[n]; ok {
		return v
	}

	fibVal := Fib(n - 1, memo) + Fib(n - 2, memo)
	memo[n] = fibVal

	return fibVal
}

func FibBig(n big.Int, memo map[string]big.Int) big.Int {
	if cmp := n.Cmp(zero); cmp == -1 || cmp == 0 {
		return *zero
	}

	if cmp := n.Cmp(two); cmp == -1 || cmp == 0 {
		return *one
	}

	if v, ok := memo[string(n.Bytes())]; ok {
		return v
	}

	subbedOne := new(big.Int).Set(&n).Sub(&n, one)
	subbedTwo := new(big.Int).Set(&n).Sub(&n, two)
	fibOne := FibBig(*subbedOne, memo)
	fibTwo := FibBig(*subbedTwo, memo)
	fibVal := new(big.Int).Add(&fibOne, &fibTwo)

	memo[string(n.Bytes())] = *fibVal

	return *fibVal
}

func FibChan(n int, res chan<- int) {
	if n < 0 {
		res <- 0
		return
	}

	if n <= 2 {
		res <- 1
		return
	}

	n1c := make(chan int)
	n2c := make(chan int)
	go FibChan(n - 1, n1c)
	go FibChan(n - 2, n2c)

	res1 := <-n1c
	res2 := <-n2c

	res <-(res1 + res2)
}
