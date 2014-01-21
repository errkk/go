// Select is for waiting on multiple channel's operations

package main

import (
    "time"
    "fmt"
)

func main() {
    // 2 unbuffered channels
    c1 := make(chan string)
    c2 := make(chan string)
    
    // Each will recieve a value after some time, like a blocking operation would
    go func() {
        time.Sleep(time.Second * 1)
        // Send to c1
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        // Send to c1
        c2 <- "two"
    }()

    // Select is in a for loop so it can run twice,
    // seems to block while it waits
    for i := 0; i < 2; i ++ {
        select {
            // getting value from ct, setting and declaring msg1 var
        case msg1 := <- c1:
            fmt.Println("Recieved:", msg1)
        case msg2 := <- c2:
            fmt.Println("Recieved:", msg2)
        }
    }

    // Total exec time is 2 seconds, as the 1 and 2 second sleeps are concurrent
}
