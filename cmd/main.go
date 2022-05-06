package main

import (
	"github.com/cjvirtucio87/fib-playground/pkg/fib"
	"log"
)

func main() {
	memo := make(map[int]int)
	expectedFib := 13
	actualFib := fib.Fib(7, memo)
	if expectedFib != actualFib {
		log.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}

	atomicMemo := fib.AtomicIntMemo{Memo: make(map[int]int)}
	actualFib = <-fib.FibChan(7, &atomicMemo)
	if expectedFib != actualFib {
		log.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}
