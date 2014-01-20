package main

import (
    "fmt"
)

func add(a int, b int) (int, string) {
    return a + b, "stringy"
}

func sum(args ...int) int {
    res := 0
    for _, v := range args {
        fmt.Println(v)
        res += v
    }
    return res
}

func intSeq(start int) func() int {
    i := start
    return func() int {
        i += 1
        return i
    }
}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    _, sres := add(2, 3)
    fmt.Println(sres)
    res := sum(1, 2, 3, 4)
    fmt.Println(res)

    array := []int{1,2,3,4,5,6,7}
    fmt.Println(sum(array...))

    fn := intSeq(20)
    fmt.Println(fn())
    fmt.Println(fn())

    i := 1
    zeroval(i)
    fmt.Println("Zeroval:", i)
    zeroptr(&i)
    fmt.Println("Zeroval:", i)
    fmt.Println("Pointer:", &i)
}
