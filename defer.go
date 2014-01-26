// Defer is used to ensure that a function call is performed later in a program’s execution

package main

import "fmt"
import "os"

func main() {
    f := createFile("/tmp/defer.txt")
    // do this later a bit like using "finally"
    defer closeFile(f)
    writeFile(f)
}

// Returns an address to os.File type
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    // Print to file
    fmt.Fprintln(f, "Stuff")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
