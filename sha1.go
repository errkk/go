package main

import (
    "fmt"
    "crypto/sha1"
)

func main() {
    s := "sha1 this string"
    h := sha1.New()

    // Write data into the hash, as bytes, so string needs to be coerced
    h.Write([]byte(s))

    // to get a byte slice, sum the hash
    // this can be for adding more stuff
    bs := h.Sum(nil)
    fmt.Printf("%x\n", h)
    fmt.Printf("%x\n", bs)
}
