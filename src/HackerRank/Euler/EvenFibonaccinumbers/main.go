package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func evenFibonacciSum(number string) {
	mNumber, _ := strconv.Atoi(number)
	sum := 0
	fib := make([]int, 0)
	fib = append(fib, 1)
	fib = append(fib, 2)
	for i := 2; fib[i-1]+fib[i-2] < mNumber; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
		fmt.Println(fib)
		if fib[i]%2 == 0 {
			sum += fib[i]
		}
	}
	fmt.Println(sum + fib[1])
}

func main() {
	isFast := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if isFast {
			isFast = false
		} else {
			evenFibonacciSum(scanner.Text())
		}
	}
}
