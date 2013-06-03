package main

import (
    "fmt"
    "runtime"
    "time"
)
func main() {
    fmt.Println(runtime.NumGoroutine())
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
    fmt.Println(runtime.NumGoroutine())
    time.Sleep(time.Second)
}
