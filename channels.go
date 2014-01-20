// This is for passing stuff between go routines
package main

import "fmt"

func main() {
    // Make a channel with the type of the thing they contain
    // Unbuffered by default, so they accept sends (chan <-) untill there's a recieve (<- chan) ready to recieve
    messages := make(chan string)

    // In a goroutine, add a thing into the channel
    // <- is channel syntax for sending stuff around
    go func() { messages <- "Oh hai"}()

    msg := <- messages
    fmt.Println(msg)


    // Buffered channels accept a limited number of strings without a corresponding reciever
    // This one can take 2 strings
    buffered_messages := make(chan string, 2)

    // Channel is buffered, so we can send these 2 things in without a corresponding recieve
    buffered_messages <- "Oh,"
    buffered_messages <- "Hai!"

    // Then we can get the stuff out later
    fmt.Println(<- buffered_messages)
    fmt.Println(<- buffered_messages)
}
