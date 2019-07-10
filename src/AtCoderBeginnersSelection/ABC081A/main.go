package main

import `fmt`

func main() {

	var s string
	fmt.Scan(&s)

	// パターン
	// 000 001 010 011 100 101 111

	if s == "000"{
		fmt.Println("0")
	}else if s == "001" || s == "010" || s == "100"{
		fmt.Println("1")
	}else if s == "011" || s == "110" || s == "101"{
		fmt.Println("2")
	}else {
		fmt.Println("3")
	}
}
