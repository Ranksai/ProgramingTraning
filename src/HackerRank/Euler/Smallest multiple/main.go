package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func smallestMultiple(number string) {
	mNumber, _ := strconv.Atoi(number)
	smallM := 2

	for i := 3; i <= mNumber; i++ {
		if smallM%i != 0 {
			remainder := smallM % i
			if i%remainder == 0 {
				smallM *= i / remainder
			} else {
				smallM *= i
			}
		}
		fmt.Println(smallM)
	}
	fmt.Println(smallM)
}

func main() {
	isFast := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if isFast {
			isFast = false
		} else {
			smallestMultiple(scanner.Text())
		}
	}
}
