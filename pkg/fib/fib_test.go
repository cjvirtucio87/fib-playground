package fib_test

import (
	"github.com/cjvirtucio87/fib-playground/pkg/fib"
	"log"
	"os"
	"testing"
)

const (
	PANIC = 0
	ERROR = 1
	WARN = 2
	INFO = 3
	DEBUG = 4
)

type Closer func()

func openLog(b *testing.B) (*os.File, Closer) {
	f, err := os.OpenFile("/tmp/test.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		b.Fatalf("failed to open test log file")
	}

	return f, func() {
		f.Close()
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
	actualFib := <-fib.FibChan(7, &atomicMemo, func(v int) {})
	expectedFib := 13
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}

func BenchmarkFibChanLarge(b *testing.B) {
	f, closer := openLog(b)
	defer closer()
	logger := log.New(f, "BenchmarkFibChanLarge", DEBUG)

	atomicMemo := fib.AtomicIntMemo{Memo: make(map[int]int)}
	actualFib := <-fib.FibChan(
		57, &atomicMemo,
		func(v int) {
			logger.Printf("value: %d", v)
		},
	)
	expectedFib := 365435296162
	if expectedFib != actualFib {
		b.Fatalf("expected [%d], got [%d]", expectedFib, actualFib)
	}
}
