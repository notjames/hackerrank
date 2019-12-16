package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func chrCount(s string, chr string) int64 {
	var matches int64
	sample := strings.Split(s, "")

	for _, c := range sample {
		if c == chr {
			matches++
		}
	}

	return matches
}

func chrCountLimit(s string, chr string, upto int64) int64 {
	var matches, i int64
	sample := strings.Split(s, "")

	for i = 0; i <= upto-1; i++ {
		if sample[i] == chr {
			matches++
		}
	}

	return matches
}

// Complete the repeatedString function below.
// matches is num_of_a's_in_s
// matches * ((n/len(s)) + (len(s) % n))
func repeatedString(s string, n int64) int64 {
	var matches, length, quotient,
		remainder, remainder_matches int64

	length = int64(len(s))
	fmt.Printf("length: %d\n", length)

	matches = chrCount(s, "a")
	fmt.Printf("matches: %d\n", matches)

	quotient = n / length
	fmt.Printf("quotient: %d\n", quotient)

	remainder = n % length
	fmt.Printf("remainder: %d\n", remainder)

	remainder_matches = chrCountLimit(s, "a", remainder)
	fmt.Printf("remainder matches: %d\n\n", remainder_matches)

	fmt.Println("equation: matches * quotient + remainder_matches = answer")
	fmt.Printf("equation: (%d * %d + %d) = %d\n", matches, quotient, remainder_matches, (matches*quotient + remainder_matches))

	return matches*quotient + remainder_matches
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
