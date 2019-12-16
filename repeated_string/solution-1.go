package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	var i, divisor, counter int64

	sample := strings.Split(s, "")
	repeated := make([]string, n+1)
	sampleLen := int64(len(sample))

	for i = 0; i <= n-1; i++ {
		if i > 0 && i%sampleLen == 0 {
			divisor++
			//fmt.Printf("i is %d -- divisor is now: %d\n", i, divisor)
		}

		if i <= sampleLen-1 {
			//fmt.Printf("i is now: %d -- length of sample is %d from s which is %q\n", i, len(sample), s)
			repeated[i] = sample[i]
			continue
		}

		repeated[i] = sample[abs(i-sampleLen)/divisor]
	}

	//fmt.Printf("size of repeated is: %d\n", len(repeated))

	for i = 0; i <= int64(len(repeated))-1; i++ {
		//fmt.Printf("  %s ", repeated[i])
		if repeated[i] == "a" {
			counter++
		}
	}

	return counter
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
