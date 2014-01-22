package main

import "fmt"
import "time"

func main() {
    // jobs channel communicates jobs from the main goroutine to the worker
    jobs := make(chan int, 5)
    done := make(chan bool)

    // This is the worker, it gets jobs from the channel
    // Its a special 2 value form of recieve
    // When the second value is false, it means the channel is closed
    go func() {
        for {
            j, more :=  <-jobs
            if more {
                fmt.Println("Recieved a job", j)
            } else {
                fmt.Println("Recieved all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 0; j <= 3; j++ {
        jobs <- j
        fmt.Println("Sent Job", j)
        time.Sleep(time.Second * 1)
    }

    close(jobs)
    fmt.Println("Sent all jobs")

    // Sync approach to block untill the worker completes by sending true to done
    <- done
}
