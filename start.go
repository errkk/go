// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"
import "time"

func main() {

    for j := 7; j <= 9; j ++ {
        fmt.Println(j)
    }

    if n := 8; n%2 == 0 {
        fmt.Println("n percents to 2", n)
    }

    hour := time.Now().Hour()

    switch {
    case hour > 12:
        fmt.Println("Its aternun")
    }

    var a[5]int
    fmt.Println(a)
    a[4] = 100
    fmt.Println(a, len(a))

    // Declare and init in one
    b := [5]int{1,2,3,4,5}
    fmt.Println("Declared", b)

    // Composing types to make a 2d array
    var twod[4][2]int
    fmt.Println(twod)
    for i := 0; i <= 3; i ++ {
        twod[i][1] = i
    }
    fmt.Println(twod)
}
