package main

import "fmt"

func main() {

    // An array, size is part of its type
    arr := [3]int{1,2,3}

    // A slice literal is declared just like an array literal, except you leave out the element count:
    sli := []string{"stuff", "fings"}
    fmt.Println(sli, arr)

    // Slices can also be made with the make builtin
    // make([]type, length, (optional) capacity) []T
    // The make function allocates an array len() and cap() builtins can inspect these properties
    s := make([]string, 4)

    fmt.Println(s)

    s[0] = "a"
    s[1] = "r"
    s[2] = "s"
    s[3] = "e"

    // append builtin can mutate slices
    s = append(s, "f", "g")
    d := s
    d[1] = "poo"
    fmt.Println(s)

    // Allocate a slice to copy the other one into
    c := make([]string, len(s))
    copy(c, s)
    c[1] = "rrr"
    fmt.Println(c)

    l := c[5:]
    fmt.Println(l)

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    mymap := make(map[string]int)
    mymap["e"] = 27
    // second return val is whether the key was found
    // Differenciates between not there and falsy values
    e, wastheree := mymap["e"]
    r, wastherer := mymap["r"]
    fmt.Println(e, wastheree)
    fmt.Println(r, wastherer)

    // ranges
    nums := []int{1,2,3}
    sum := 0
    fmt.Println(nums)
    for _, i := range(nums) {
        sum += i
    }
    fmt.Println(sum)

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
