package fib_test

import (
	"github.com/cjvirtucio87/fib-playground/pkg/fib"
	"math/big"
	"testing"
)

func BenchmarkFibBig(b *testing.B) {
	memo := make(map[string]big.Int)
	expectedFib := big.NewInt(13)
	actualFib := fib.FibBig(*big.NewInt(7), memo)
	if cmp := expectedFib.Cmp(&actualFib); cmp != 0 {
		b.Fatalf("expected [%v], got [%v]", expectedFib, actualFib)
	}
}

func BenchmarkFibSmall(b *testing.B) {
	memo := make(map[int]int)
	expectedFib := 13
	actualFib := fib.Fib(7, memo)
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}

func BenchmarkFibChanSmall(b *testing.B) {
	actualChan := make(chan int)
	go fib.FibChan(7, actualChan)
	expectedFib := 13
	actualFib := <-actualChan
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}
