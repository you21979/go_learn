package main

import (
    "fmt"
    "runtime"
    "time"
)
func main() {
    go func() {
        fmt.Println("end")
    }()
    go func() {
        fmt.Println("return")
        return
    }()
    go func() {
        fmt.Println("exit")
        runtime.Goexit()
    }()
    fmt.Println("wait")
    time.Sleep(time.Second)
}
