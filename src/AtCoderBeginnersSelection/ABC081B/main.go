package main

import `fmt`

func main() {

	var n int
	count := 0
	isEven := true
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	for isEven{
		for i, num := range nums {
			if num%2 != 0 {
				fmt.Println(count)
				isEven = false
				break
			}
			nums[i] = num / 2
		}
		count++
	}
}
