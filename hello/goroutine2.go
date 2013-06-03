package main

import (
    "fmt"
    "runtime"
    "time"
)
func kick(i int){
    go func() {
        time.Sleep(time.Second)
        fmt.Println(i)
    }()
}
func main() {
    fmt.Println(runtime.NumGoroutine())
    for i := 0; i < 100000; i++ {
        kick(i)
    }
    fmt.Println("wait")
    fmt.Println(runtime.NumGoroutine())
    time.Sleep(time.Second * 3)
}
