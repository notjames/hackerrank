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

func isSigned(n int32) bool {
	return math.Signbit(float64(n))
}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var lastHere, here, valleyCount int32
	var isValley bool

	fmt.Printf("s is: %+v\n", s)
	for _, step := range strings.Split(s, "") {
		fmt.Printf("step is: %s\n", step)

		if string(step) == "U" {
			here++
		}
		if string(step) == "D" {
			here--
		}

		isValley = isSigned(here)
		fmt.Printf("lastHere is: %d    here is: %d   isSigned is: %b\n", lastHere, here, isValley)

		if lastHere == 0 && isValley == true {
			fmt.Printf("  valley_count++\n")
			valleyCount++
		}

		fmt.Printf("valleyCount is now: %d\n", valleyCount)
		lastHere = here
	}

	return valleyCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

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
