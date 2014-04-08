package main

import "fmt"

type Any interface{}

type State interface {
    Begin(a Any)
    End(a Any)
    Update(a Any)
}

type Vector struct {
    x float64
    y float64
    z float64
}

type MapObject struct {
    Vector
    dir uint16
}

const (
    state_init = iota
    state_battle
    state_sizeof
)

type Monster struct {
    MapObject
    state int
    id int
}

type StateInit struct {
}
func (s *StateInit) Begin(a Any) {
}
func (s *StateInit) Update(a Any) {
    fmt.Println("init mode")
}
func (s *StateInit) End(a Any) {
}

type StateBattle struct {
}
func (s *StateBattle) Begin(a Any) {
}
func (s *StateBattle) Update(a Any) {
    fmt.Println("battle mode")
}
func (s *StateBattle) End(a Any) {
}

func run(s []State, m Monster) {
    s[m.state].Begin(&m)
    s[m.state].Update(&m)
    s[m.state].End(&m)
}

func main() {
    var state []State = []State{&StateInit{},&StateBattle{}}
    var m Monster = Monster{}
    m.state = state_init
    run(state, m)
    m.state = state_battle
    run(state, m)
}
