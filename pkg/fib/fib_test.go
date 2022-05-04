package fib_test

import (
	"github.com/cjvirtucio87/fib-playground/pkg/fib"
	"testing"
)

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
