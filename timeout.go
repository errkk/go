// Using this select timeout pattern requires communicating results over channels.
// This is a good idea in general because other important Go features are based on channels and select.

package main

import (
    "fmt"
    "time"
)

func main() {

    // Make a buffered channel
    c1 := make(chan string, 1)

    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    // Use a single select to see what's in the channel
    // It takes whatever comes first, so if the timeout yields first then that case will be satisfied
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    // Trying again with a function that will return before the timeout
    c2 := make(chan string, 1)

    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
