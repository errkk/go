package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // Range over the elements in the channel
    // Terminates after 2 because its now closed
    for elem := range queue {
        fmt.Println(elem)
    }

    // If queue hadn't been closed, a deadlock runtime error would happen
}

