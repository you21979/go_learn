package main
import "fmt"

type Hoge struct {
    aaa int
    bbb string
}

func main(){
    hoge := Hoge{100, "aaa"}
    phoge := new(Hoge)
    fmt.Println(hoge.bbb)
    fmt.Println(phoge.aaa)
}
