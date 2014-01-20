package main

import "fmt"

type person struct {
    name string
    age  int
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})

    // Yields a pointer to the struct
    fmt.Println(&person{name: "Ann", age: 40})

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // dot on a pointer
    sp := &s
    fmt.Println(sp.age)

    // Mutable
    sp.age = 45
    fmt.Println(sp.age)

    s.age = 53
    fmt.Println(s)

}
