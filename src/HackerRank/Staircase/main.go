package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	stairString := ""

	for i := 1; i <= int(n); i++ {
		spaceString := ""
		stairString += "#"
		for j := 0; j < int(n)-i; j++ {
			spaceString += " "
		}
		fmt.Printf("%s\n", spaceString+stairString)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
