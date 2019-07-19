package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var min int32 = math.MaxInt32
	var max int32 = 0
	var minSum int64
	var maxSum int64

	for i := range arr {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	if min == max {
		fmt.Println(int(min)*(len(arr)-1), int(min)*(len(arr)-1))
		return
	}
	for i := range arr {
		if arr[i] == min {
			minSum += int64(arr[i])
		} else if arr[i] == max {
			maxSum += int64(arr[i])
		} else {
			minSum += int64(arr[i])
			maxSum += int64(arr[i])
		}
	}

	fmt.Println(minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
