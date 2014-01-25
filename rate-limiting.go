package main

import (
    "fmt"
    "time"
)

func main() {
    // Ratelimiting incoming requests
    // Serving them off a channel
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // limiter is a channel
    // It recieves a tick every 200ms
    limiter := time.Tick(time.Millisecond * 200)

    for req := range requests {
        <- limiter
        fmt.Println("Request", req, time.Now())
    }

    // To allow bursts, buffering the limiter channel
    // This allowes bursts of up to 3 events
    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // Putting a request to the bursty limiter every 200 ms
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    // 5 requests for the bursty limiter, the first 3 should benefit from the burstableness
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)

    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
