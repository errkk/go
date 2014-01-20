// In Go it’s idiomatic to communicate errors via an explicit, separate return value.

package main

import (
    "errors"
    "fmt"
)


// By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
    if arg != 41 {
        return -1, errors.New("Nope, thats not right")
    }
    // Nil value in error position indicates no error
    return arg + 3, nil
}

// Custom type to explicitly represent an argument error.
type argError struct {
    arg int
    prob string
}

// Error reciever
func (e *argError) Error() string {
    return fmt.Sprintf("%d, %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        // Make a reference to a new struct of argError type
        return -1, &argError{arg, "Nope"}
    }
    return arg + 3, nil
}

func main() {
    fmt.Println(f1(41))
    fmt.Println(f1(42))

    fmt.Println(f2(41))
    fmt.Println(f2(42))

    // Trying with some inline error checking in the if line (ideomatic)
    for _, i := range []int{7, 42} {
        // Set the vars for the scop of the if
        // check the nilness of e
        if r, e := f2(i); e != nil {
            fmt.Println(i, "got an error", e)
        } else {
            fmt.Println(i, "worked", r)
        }
    }

    // To programatically use the data from a custom error,
    // you get the error as an instance of the custom error type
    // via type assertion
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
