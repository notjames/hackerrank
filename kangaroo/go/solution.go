package main

import (
	"fmt"
)

// UPPERBOUND just a simple const
const UPPERBOUND int = 10e5

func kangaroo(k1s, k1d, k2s, k2d int) string {
	fmt.Printf("k1s and k2s are respectively: %v and %v\n", k1s, k1d)
	fmt.Printf("k2s and k2d are respectively: %v and %v\n\n", k2s, k2d)
	for j := 1; j <= UPPERBOUND; j++ {
    if ( (k1d * j + k1s) == (k2d * j + k2s) ) {
      return "YES"
    }
	}

	return "NO"
}

func main() {
  var k1s, k2s, k1d, k2d int
  fmt.Scan(&k1s, &k1d, &k2s, &k2d)

	result := kangaroo(k1s, k1d, k2s, k2d)

	fmt.Println(result)
}
