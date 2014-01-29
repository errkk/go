package main

import "fmt"
import "strconv"

func main() {
    // 64 tells how many bits of precision to parse.
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // Atoi is for parsing integers
    k, _ := strconv.Atoi("135")
    fmt.Println(k)
}
