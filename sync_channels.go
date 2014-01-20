// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.

package main

import (
    "fmt"
    "time"
)

// This is the function to run in the goroutine,
// It takes a channel and sends to it when its done
// That will notify another goroutine
func worker(done chan bool) {
    fmt.Println("Working")
    time.Sleep(time.Second)
    fmt.Println("Done")

    // Send to the channel
    done <- true
}

// You can specify the direction of channels to increase type safety

// ping only accepts a channel for sending values
// Compile error if a recieving channel was used
func ping(pings chan<- string, msg string) {
    fmt.Println("ping")
    pings <- msg
}

// accepts a pings channel for recieving and a pongs one for sending
func pong(pings <-chan string, pongs chan<- string) {
    fmt.Println("pong")
    // Recieves from pings, and saves to msg var
    msg := <- pings
    // sends msg var out to pongs
    pongs <- msg
}

func main() {

    // Make a channel to link shit up
    done := make(chan bool, 1)
    go worker(done)

    // This just blocks until notifcation is recieved from the worker in the other gr
    <- done

    // Making some single item buffered channels
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    // Ping function and pings channel for sending a message (passed to the func)
    ping(pings, "passed message")

    // Takes the reciever and the sender, sends back out what it recieves
    pong(pings, pongs)

    // Now the message is recieved by pongs
    fmt.Println(<-pongs)
}
