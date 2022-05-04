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

func BenchmarkFibLarge(b *testing.B) {
	memo := make(map[int]int)
	expectedFib := 365435296162
	actualFib := fib.Fib(57, memo)
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}

func BenchmarkFibChanSmall(b *testing.B) {
	atomicMemo := fib.AtomicIntMemo{Memo: make(map[int]int)}
	actualChan := make(chan int)
	go fib.FibChan(7, actualChan, &atomicMemo)
	expectedFib := 13
	actualFib := <- actualChan
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}

func BenchmarkFibChanLarge(b *testing.B) {
	atomicMemo := fib.AtomicIntMemo{Memo: make(map[int]int)}
	actualChan := make(chan int)
	go fib.FibChan(57, actualChan, &atomicMemo)
	expectedFib := 365435296162
	actualFib := <- actualChan
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}
