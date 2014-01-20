package main

import "fmt"

type rect struct {
    width, height int
}

// Reciever for the rect type.
// This recieves the rect as a pointer, so can mutate it
// Using pointer reciever type also avoids copying on method calls
func (r *rect) area() int {
    res := r.width * r.height
    // Mutate Struct
    r.width = r.width * 2
    return res
}

// Value reciever for rect
// Any changes to rect inside here aren't appied to the struct
func (r rect) perim() int {
    res := 2 * r.width  + 2 * r.height
    // This mutation doesnt go anywhere
    r.width = 2000
    return res
}

func main() {
    my_instance := rect{width: 10, height: 20}

    fmt.Println("area:", my_instance.area())
    fmt.Println("perim", my_instance.perim())

    // Method calls on a pointer to the instance
    pointer := &my_instance
    fmt.Println("pointer area:", pointer.area())
    fmt.Println("pointer area:", pointer.area())

    fmt.Println(pointer, my_instance)
}
