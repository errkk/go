package main

import "fmt"

func f(from string) {
    for i := 0; i < 4; i ++ {
        fmt.Println(from, ":", i)
    }
}
func main() {
    f("Direct")

    // This makes it run in a goroutine
    go f("Go")

    // Ana anon functs
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    var input string
    fmt.Scanln(&input)
    fmt.Println("Done")
}
