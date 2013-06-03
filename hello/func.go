package main

import "fmt"

func hoge(param string) string{
    return param;
}

func fuge() string{
    return "Hello World";
}

func main() {
    fmt.Println(hoge(fuge()))
}
