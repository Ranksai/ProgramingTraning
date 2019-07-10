package main

import `fmt`

func checkNumber(number int, a int, b int)  bool{
	number2 := number/10
	number3 := number2/10
	number4 := number3/10
	number1 := number - (number4*1000 +number3*100 +number2*10)

	if number1+number2+number3+number4 >= a &&  number1+number2+number3+number4 <= b {
		return true
	}
	return false
}
func main() {
	var (
		n int
		a int
		b int
	)
	sum := 0
	fmt.Scan(&n, &a, &b)

	for i := 0; i <= n; i++{
		if checkNumber(i, a, b) {
			sum += i
		}
	}
	fmt.Println(sum)
}
