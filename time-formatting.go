package main

import (
    "fmt"
    "time"
)


func main() {
    p := fmt.Println
    t := time.Now()
    p(t.Format("2006-01-02T15:04:05Z07:00"))

    withNanos := "2006-01-02T15:04:05.999999999-07:00"
    t1, e := time.Parse(
        withNanos,
        "2012-11-01T22:08:41.117442+00:00")
    p(t1)
    kitchen := "3:04PM"
    t2, e := time.Parse(kitchen, "8:41PM")
    p(t2)
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
