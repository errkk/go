package main

import (
    "fmt"
    "time"
)

func main() {
    // Timers will provide a channel that will yield something after the delay
    timer1 := time.NewTimer(time.Second * 2)

    // Blocking channel syntax
    <- timer1.C
    fmt.Println("Timer 1 expired")

    // A blocking timer on a channel is like time.Sleep
    // But a channel can be canceled

    timer2 := time.NewTimer(time.Second * 3)

    go func() {
        <- timer2.C
        fmt.Println("Timer 2 expired")
    }()

    stop := timer2.Stop()
    if stop {
        fmt.Println("Timer 2 stopped")
    }

    // Tickers are channels that keep giving out ticks
    ticker := time.NewTicker(time.Millisecond * 500)

    go func() {
        for i := range ticker.C {
            fmt.Println("Tick", i)
        }
    }()

    // The other way to wait
    time.Sleep(time.Second * 3)
    ticker.Stop()
    fmt.Println("Ticker Stopped")
}
