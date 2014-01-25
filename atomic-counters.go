package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {
    var ops uint64 = 0

    // TO simulate concurrency starting 50 goroutines
    for i := 0; i < 50; i++ {
        // Each routine increments the counter about every millisecond
        go func() {
            for {
                // Passing the memory address of the counter var
                // AddUint64 automatically adds to this var
                atomic.AddUint64(&ops, 1)
                // Allow other goroutines to proceed
                runtime.Gosched()
            }
        }()
    }

    // Wait for a second for it to do stuff
    time.Sleep(time.Second * 1)
    // To safely load the value while its being updated by other goroutines
    // Use LoadUint64 to take the memory reference and copy out the value
    opsFinal := atomic.LoadUint64(&ops)

    fmt.Println(opsFinal)
}
