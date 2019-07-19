package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func largestPrimeFactor(number string) {
	mNunber, _ := strconv.ParseInt(number, 10, 64)

	for mNunber%2 == 0 {
		mNunber = mNunber / 2
	}
	max := 2
	for i := 3; mNunber != 1; i += 2 {
		for mNunber%int64(i) == 0 {
			mNunber = mNunber / int64(i)
		}
		max = i
	}
	fmt.Println(max)
}

func main() {
	isFast := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if isFast {
			isFast = false
		} else {
			largestPrimeFactor(scanner.Text())
		}
	}
}
