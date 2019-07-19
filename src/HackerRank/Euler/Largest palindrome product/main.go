package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPalindrome(i, j int) bool {

	pal := strings.Split(strconv.Itoa(i*j), "")
	for l := 0; l < 3; l++ {
		if pal[l] != pal[5-l] {
			return false
		}
	}
	return true
}

func largestPalindromeProduct(number string) {
	mNumber, _ := strconv.Atoi(number)
	max := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if i < j || i*j >= mNumber {
				break
			} else if i*j > max && i*j > 100000 && isPalindrome(i, j) {
				//fmt.Println(i, j, i*j)
				max = i * j
			}
		}
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
			largestPalindromeProduct(scanner.Text())
		}
	}
}
