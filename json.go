package main

import (
    "os"
    "fmt"
    "encoding/json"
)

// Structs for demonstrating encoding and decoding custom types
type Response1 struct {
    Page int
    Fruits []string
}

type Response2 struct {
    Page int `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    bolB, _ := json.Marshal(true)
    fmt.Println(bolB)

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // JSON can export custom types
    // It will only include exported (caps) fields on the encoded output
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // You can give custom key names for json in the type by using tags
    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // Decoding ____
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // Make a map for the decoded data to go in
    // It will hold a map of strings to arbitrary data types
    var dat map[string]interface{}

    // This does the decoding. It might have an error
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // To use the decoded data, it needs to be cast into suitable types
    num := dat["num"].(float64)
    fmt.Println("Number", num)

    // Accessing nested data requires a series of casts
    // This casts it into a slice of interfaces
    strs := dat["strs"].([]interface{})

    // Accessing and casting the first string
    str1 := strs[0].(string)
    fmt.Println(str1)

    // Alternatively. Data can be decoded into a custom type
    // This adds additional type safety without having to manually assert types
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    // & is there so when res is accessed, its a pointer
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // Above, bytes and strings were used to display json data on stdOut via println
    // JSON encoder can also be pointed at a stream, like a file or stdOut
    enc := json.NewEncoder(os.Stdout)
    // Some data
    d := map[string]int{"apple": 5, "lettuce": 7}
    // and that goes straight to the stream
    enc.Encode(d)
}
