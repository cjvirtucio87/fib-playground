package fib

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