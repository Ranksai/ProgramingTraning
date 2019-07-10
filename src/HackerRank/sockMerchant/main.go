package sockMerchant

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func unset(s []int32, i int) []int32 {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	fmt.Println(ar)
	count := 0

	for i := 0; int32(i) < n; i++{
		for j := i+1; int32(j) < n; j++{
			if (len(ar) >= j+1) && ar[i] == ar[j] {
				count ++
				ar = unset(ar, j)
				break
			}
		}
	}
	return int32(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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

