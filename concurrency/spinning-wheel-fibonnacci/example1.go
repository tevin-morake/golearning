package main

import (
	"fmt"
	"time"
)

/*
 * Forgive my many comments but I hope you find it worthwile.
 * The following code creates a spinner like effect in a terminal shell while trying to calculate the fibonnacci sequence of a number n using the recursive algorithm
 * The spinner function is launched into it's own goroutine and runs concurrently(at the same time) as the main goroutine(the execution of the main function)
 * Once the fib(...) function is done running, and prints to stdout, the main goroutine exits. All other existing goroutines also exit
 */

func main() {
	go spinner(100 * time.Millisecond) // go keyword launches spinnner function into a goroutine of it's own
	const n = 45
	fmt.Println("Main Goroutine is running")

	fibN := fib(n) //slow

	fmt.Printf("\rFibonnaci(%d) = %d\n", n, fibN)
	fmt.Println("All goroutines exiting")
}

func spinner(delay time.Duration) {
	fmt.Println("Second goroutine is running")
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
