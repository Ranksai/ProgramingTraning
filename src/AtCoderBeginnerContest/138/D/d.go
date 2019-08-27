package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dfs(start, before int64, arr [][]int64, point []int64) {
	for i := range arr[start] {
		if arr[start][i] == before{
			continue
		}
		point[arr[start][i]] += point[start]
		dfs(int64(arr[start][i]), start, arr, point)
	}
}

func calc(arr [][]int64, point []int64, n int , q int) {
	dfs(0, -1, arr, point)
	for i := range point{
		fmt.Print(point[i])
		fmt.Print(" ")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	strArr := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(strArr[0])
	q, _ := strconv.Atoi(strArr[1])
	arr := make([][]int64, n)
	point := make([]int64, n)

	for i := 0; i< n-1; i++ {
		strArr = strings.Split(readLine(reader), " ")
		tmpA, _ := strconv.ParseInt(strArr[0],10,64)
		tmpB, _ := strconv.ParseInt(strArr[1],10,64)
		arr[tmpA-1] = append(arr[tmpA-1], tmpB-1)
		arr[tmpB-1] = append(arr[tmpB-1], tmpA-1)
	}
	for i := 0; i< q; i++ {
		strArr = strings.Split(readLine(reader), " ")
		tmpP, _ := strconv.Atoi(strArr[0])
		tmpX, _ := strconv.ParseInt(strArr[1],10,64)
		point[tmpP-1] += tmpX
	}
	calc(arr, point, n, q)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
