package main

import "fmt"
import "sort"

// Custom sorting
type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    strs := []string{"c", "a", "b"}
    // Sorting mutates original and doesnt return a copy
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)

    // Custom sort
    // By following this same pattern of creating a custom type,
    // implementing the three Interface methods on that type,
    // and then calling sort.Sort on a collection of that custom type,
    // we can sort Go slices by arbitrary functions.
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
