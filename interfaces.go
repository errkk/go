package main

import "fmt"
import "math"

// Geometry type, anything with area and perim recievers
type geometry interface {
    area() float64
    perim() float64
}

type square struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (s square) area() float64 {
    return s.width * s.height
}

func (s square) perim() float64 {
    return 2 * s.height + 2 * s.width
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2* math.Pi * c.radius
}

// Measure takes geometry type
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := square{width: 10, height: 10}
    c := circle{radius: 5}

    measure(s)
    measure(c)
}
