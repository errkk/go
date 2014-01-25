// Mutex is for safely managing more complex [than atomic] state
// accross goroutines
package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {
    // State is a map (of ints with int keys)
    var state = make(map[int]int)

    var mutex = &sync.Mutex{}

    // Ops will count how many operations are performed
    var ops int64 = 0


    // Making 100 goroutines to repeatedly read against the state
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                // Locking the mutex to get exclusive access to the state
                mutexLock()
                // Increment total with what's in the state at this key
                total += state[key]
                // Unlock mutex and increpent ops count (same as in atomic-counters)
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                // In order to ensure that this goroutine doesn’t starve the scheduler,
                // we explicitly yield after each operation with runtime.Gosched()
                runtime.Gosched()
            }
        }()
    }

    // Also starting 10 goroutines to simulate writes
    // Using the same pattern as for reads
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    // Give them both a second to do their shit
    time.Sleep(time.Second)

    // Safely load the total opscount from atomic
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    // Lock and unlock mutex to see what's in the state
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()

}
