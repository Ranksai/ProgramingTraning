package main

import `fmt`

func main()  {

	a := 0
	b := 0
	c := 0
	x := 0
	count := 0

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if 500*i + 100*j + 50*k == x {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}