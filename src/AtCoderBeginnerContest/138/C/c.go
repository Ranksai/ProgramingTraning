package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)
func avg (a, b float64) float64 {
	return (a+b)/2.0
}
func calc(arr []float64, times int) {
	sort.Float64s(arr)
	tmpAvg := arr[0]
	for i := 1; i < len(arr); i++ {
		tmpAvg = avg(tmpAvg,arr[i])
	}
	fmt.Println(tmpAvg)
}

func main() {
	arr := make([]float64, 0)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	times, _:= strconv.Atoi(readLine(reader))
	strArr := strings.Split(readLine(reader), " ")
	for i := 0; i< times ; i++ {
		tmp, _ := strconv.ParseFloat(strArr[i],64)
		arr = append(arr, tmp)
	}
	calc(arr, times)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}