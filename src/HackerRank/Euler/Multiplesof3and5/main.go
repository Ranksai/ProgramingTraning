package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func multiples(naturalNumbersBelow string) {
	number, _ := strconv.Atoi(naturalNumbersBelow)
	sum := 0
	number--
	sum3 := number / 3
	sum5 := number / 5
	sum15 := number / 15
	sum = (1+sum3)*3*sum3/2 + (1+sum5)*5*sum5/2 - (1+sum15)*sum15*15/2
	fmt.Println(sum)
}

func main() {
	isFast := true
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if isFast {
			isFast = false
		} else {
			multiples(scanner.Text())
		}
	}
}
