package fib

import (
	"sync"
)

type AtomicIntMemo struct {
	m sync.RWMutex
	Memo map[int]int
}

func (am *AtomicIntMemo) Get(k int) (int, bool) {
	am.m.RLock()
	defer am.m.RUnlock()
	var res int
	var ok bool
	if v, ok := am.Memo[k]; ok {
		res = v
	}

	return res, ok
}

func (am *AtomicIntMemo) Put(k int, v int) {
	am.m.Lock()
	defer am.m.Unlock()
	am.Memo[k] = v
}

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

func FibChan(n int, res chan<- int, memo *AtomicIntMemo) {
	defer close(res)
	if n < 0 {
		res <- 0
		return
	}

	if n <= 2 {
		res <- 1
		return
	}

	if v, ok := memo.Get(n); ok {
		res <- v
		return
	}

	n1c := make(chan int)
	n2c := make(chan int)
	go FibChan(n - 1, n1c, memo)
	go FibChan(n - 2, n2c, memo)

	res1 := <-n1c
	res2 := <-n2c

	resVal := res1 + res2
	memo.Put(n, resVal)

	res <- resVal
}
