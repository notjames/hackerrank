// https://www.hackerrank.com/challenges/apple-and-orange/problem
package main

import (
        "bufio"
        "fmt"
        "os"
        "strings"
        "strconv"
        "errors"
)

const upperbound int = 10e5
const lowerbound int = -10e5
const expectInpLines int = 5

type error interface {
  Error() string
}

func getContents() (contents []string) {
  reader := bufio.NewScanner(os.Stdin)
  buf := make([]byte, 0, 64*1024)
  reader.Buffer(buf, 1024*1024)

  for reader.Scan() {
    line := reader.Text()

    // fmt.Printf("  ** On input line is: %v\n", line)
    line = strings.TrimSuffix(line, "\n")
    // fmt.Printf("  ** After trimming line is: %v\n", line)

    contents = append(contents, line)
  }

  return contents
}

func cleanStrArr(arr *[]string) ([]string) {
  clean := []string{}

  for _, element := range *arr {
    if element != "" {
      clean = append(clean, element)
    }
  }

  return clean
}

func arrayConvAtoi(arr *[]string) ([]int) {
  intsarr := []int{}

  for _, element := range *arr {
    converted, _ := strconv.Atoi(element)
    intsarr = append(intsarr, converted)
  }

  return intsarr
}

// s: house right side
// t: house left side
// a: larry is in apple tree (left side)
// b: rob is in orange tree (right side)
// apples: apples thrown (by Larry)
// oranges: oranges thrown (by Rob)
func countApplesAndOranges(s, t, a, b int, apples, oranges []int) error {
  // fmt.Printf("s: %v, t: %v, a: %v, b: %v, apples: %v, oranges: %v\n", s, t, a, b, apples, oranges)

  //m := len(apples)
  //n := len(oranges)
  var larry, rob int

  if s < 1 || s > upperbound {
    msg := fmt.Sprintf("left side of house (%v) violates boundary constraint (lower: 1, upper: %v).\n", s, upperbound)
    return errors.New(msg)
  }

  if t < 1 || t > upperbound {
    msg := fmt.Sprintf("right side of house (%v) violates boundary constraint (lower: 1, upper: %v).\n", t, upperbound)
    return errors.New(msg)
  }

  for _, distance := range oranges {
    if lowerbound <= a || a <= upperbound {
      poss := a + distance

      // fmt.Printf("(Rob) Does (%d + %d = %d) fall within %v and %v? ", a, distance, poss, s, t)
      if poss >= s && poss <= t {
        // fmt.Println(" yes")
        // fmt.Printf("(%d + %d = %d) larry gets a point because %v falls within %v and %v\n", a, distance, poss, poss, s, t)
        larry++
      } else {
        // fmt.Println(" no")
      }
    } else {
      msg := fmt.Sprintf("distance boundary (%v) violates constraint (lower: %v, upper: %v).\n", lowerbound, upperbound)
      return errors.New(msg)
    }
  }

  for _, distance := range apples {
    if lowerbound <= b || b <= upperbound {
      poss := b + distance

      // fmt.Printf("(Larry) Does (%d + %d = %d) fall within %v and %v? ", b, distance, poss, s, t)
      if poss >= s && poss <= t {
        // fmt.Println(" yes")
        // fmt.Printf("(%d + %d = %d) rob gets a point because %v falls within %v and %v\n", b, distance, poss, poss, s, t)
        rob++
      } else {
        // fmt.Println(" no")
      }
    } else {
      msg := fmt.Sprintf("distance boundary (%v) violates constraint (lower: %v, upper: %v).\n", lowerbound, upperbound)
      return errors.New(msg)
    }
  }

  fmt.Printf("%d\n%d\n", larry, rob)
  return nil
}

func main() {
  var z, x, c, v string
  var contents, fruita, fruitb []string
  contents = getContents()

  var splitv[] string

  splitv = strings.Split(contents[0], " ")
  splitv = cleanStrArr(&splitv)
  z, x = splitv[0], splitv[1]
  s, _ := strconv.Atoi(z)

  // larry's throw and rob's throw
  splitv = strings.Split(contents[1], " ")
  c, v = splitv[0], splitv[1]
  t, _ := strconv.Atoi(x)

  // apples
  splitv = strings.Split(contents[3], " ")
  fruita = splitv
  a, _ := strconv.Atoi(c)
  apples := arrayConvAtoi(&fruita)

  // oranges
  splitv = strings.Split(contents[4], " ")
  fruitb = splitv
  b, _ := strconv.Atoi(v)
  oranges := arrayConvAtoi(&fruitb)

  err := countApplesAndOranges(s, t, a, b, oranges, apples)

  if err != nil {
    fmt.Print(err)
    os.Exit(10)
  }
}
