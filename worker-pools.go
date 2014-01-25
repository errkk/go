package main

import (
    "fmt"
    "time"
)

// Worker will process from the jobs channel and send the results
// into the results channel. There will be multiple instances of it
// running concurrently
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("Worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Instanciating lots of workers,
    // Each connected to the same channels
    // This will be blocked initially, as there are no jobs
    for w := 1; w <= 9; w ++ {
        go worker(w, jobs, results)
    }

    // Put some jobs into the channel
    // Then close it to indicate that's all the jobs there are
    for j := 1; j <= 29; j ++ {
        jobs <- j
    }
    close(jobs)

    // Collect the results
    for a := 1; a <= 19; a ++ {
        <- results
    }
}
