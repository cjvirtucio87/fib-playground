package main

import (
	"log"
)

func fib(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	if v, ok := memo[n]; ok {
		return v
	}

	fibVal := fib(n - 1, memo) + fib(n - 2, memo)
	memo[n] = fibVal

	return fibVal
}

func fibChan(n int, res chan<- int) {
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
	go fibChan(n - 1, n1c)
	go fibChan(n - 2, n2c)

	res1 := <-n1c
	res2 := <-n2c

	res <-(res1 + res2)
}

func main() {
	memo := make(map[int]int)
	expectedFib := 13
	actualFib := fib(7, memo)
	if expectedFib != actualFib {
		log.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}

	actualChan := make(chan int)
	go fibChan(7, actualChan)
	actualFib = <-actualChan
	if expectedFib != actualFib {
		log.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}
