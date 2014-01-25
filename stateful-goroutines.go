// Alternative approach to using mutexs.
// This is a little more complicated but in some cases would
// be more reliable than multiple mutexs if other channels are involved
package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// State will be handled by a single goroutine
// This guarantees that data can't be corrupted by concurrent access
// To read or write state, goroutines send messages to the owning goroutine

// Types for the read and write messages
type readOp struct {
    key int
    resp chan int
}

type writeOp struct {
    key int
    val int
    resp chan bool
}

func main() {
    // Counting how many operations
    var ops int64 = 0

    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    // Goroutine repeatedly selects on the reads and writes channels
    // responding to requests as they arrive
    // response is executed by first performing the requested action
    // then sending a value on the response channel
    // Its a success flag for writes, and the value in the case of reads
    go func() {
        // State is private to this goroutine
        var state = make(map[int]int)
        for {
            select {
            case read := <- reads:
                read.resp <- state[read.key]
            case write := <- writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // 100 read requests
    for r := 0; r < 100; r++ {
        go func() {
            for {
                // Each read requires constructing a readOp, sending it over the reads channel,
                // and the receiving the result over the provided resp channel.
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // 10 writing routines
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    time.Sleep(time.Second)
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
