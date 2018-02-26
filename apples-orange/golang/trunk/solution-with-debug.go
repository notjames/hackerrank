package debug

import (
        "bufio"
        "fmt"
        "os"
        "strings"
        "strconv"
)

const upperbound int = 10e5
const lowerbound int = -10e5
const expectInpLines int = 5

func getLine() (line string) {
        reader := bufio.NewReader(os.Stdin)

        line, _ = reader.ReadString('\n')
        line = strings.TrimSuffix(line, "\n")

        return line
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
func countApplesAndOranges(s, t, a, b int, apples, oranges []int) {
  fmt.Printf("s: %v, t: %v, a: %v, b: %v, apples: %v, oranges: %v\n", s, t, a, b, apples, oranges)
  return
}

func main() {
        var line string
        var z, x, c, v string
        var fruita, fruitb []string

        for i := 0; i < expectInpLines; i++ {
                var splitv[] string
                line = getLine()

                switch i {
                // house right and left
                case 0:
                        splitv = strings.Split(line, " ")
                        splitv = cleanStrArr(&splitv)
                        fmt.Printf("case %v: line is: %T, %v\n", i, line, line)
                        fmt.Printf("splitv is: %T, %v, with count %d elements\n", splitv, splitv, len(splitv))
                        z, x = splitv[0], splitv[1]
                // larry's throw and rob's throw
                case 1:
                        splitv = strings.Split(line, " ")
                        c, v = splitv[0], splitv[1]
                        fmt.Printf("case %v: line is: %T, %v\n", i, line, line)
                        fmt.Printf("splitv is: %T, %v, with count %d elements\n", splitv, splitv, len(splitv))
                case 2:
                        // apples
                        splitv = strings.Split(line, " ")
                        fruita = splitv
                        fmt.Printf("case %v: line is: %T, %v\n", i, line, line)
                        fmt.Printf("fruita is: %T, %v, with count %d elements\n", fruita, fruita, len(fruita))
                case 3:
                        // oranges
                        splitv = strings.Split(line, " ")
                        fruitb = splitv
                        fmt.Printf("case %v: line is: %T, %v\n", i, line, line)
                        fmt.Printf("fruitb is: %T, %v, with count %d elements\n", fruitb, fruitb, len(fruitb))

                }

                if i <= (expectInpLines-1) && line == "" {
                        i--
                }
        }


        s, _ := strconv.Atoi(z)
        t, _ := strconv.Atoi(x)
        a, _ := strconv.Atoi(c)
        b, _ := strconv.Atoi(v)
        apples := arrayConvAtoi(&fruita)
        oranges := arrayConvAtoi(&fruitb)

        countApplesAndOranges(s, t, a, b, oranges, apples)
}
