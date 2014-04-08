package main

import "fmt"

type Any interface{}

type State interface {
    Begin(a Any)
    End(a Any)
    Update(a Any)
}

type Vector2 struct {
    x float64
    y float64
}
type Vector3 struct {
    Vector2
    z float64
}

type MapObject struct {
    Vector3
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

type StateInit struct {}
func (s *StateInit) Begin(a Any) {
    fmt.Println("init begin")
}
func (s *StateInit) Update(a Any) {
    fmt.Println("init update")
}
func (s *StateInit) End(a Any) {
    fmt.Println("init end")
}

type StateBattle struct {}
func (s *StateBattle) Begin(a Any) {
    fmt.Println("battle begin")
}
func (s *StateBattle) Update(a Any) {
    fmt.Println("battle update")
}
func (s *StateBattle) End(a Any) {
    m := a.(*Monster)
    m.x = 100
    m.z = 101
    fmt.Println(a)
    fmt.Println("battle end")
}

type StateMachine struct {
    current State
}
func (sm *StateMachine) GetState() State {
    return sm.current
}
func (sm *StateMachine) _SetState(state State) {
    sm.current = state 
}
func (sm *StateMachine) SetState(state State, a Any) {
    if sm.current != nil {
        sm.current.End(a);
    }
    sm._SetState(state)
    if sm.current != nil {
        sm.current.Begin(a);
    }
}
func (sm *StateMachine) Update(a Any) {
    if sm.current != nil {
        sm.current.Update(a);
    }
}

type StateManager struct {
    sm StateMachine
    state int
    state_table []State
}
func (m *StateManager) Add(idx int, state State) {
    m.state_table[idx] = state
}
func (m *StateManager) Update(a Any) {
    m.sm.Update(a)
}
func (m *StateManager) GetState() int {
    return m.state
}
func (m *StateManager) SetState(idx int, a Any) {
    m.state = idx
    m.sm.SetState(m.state_table[idx], a)
}
func NewStateManager(max int) *StateManager {
    m := new(StateManager)
    m.state = 0
    m.sm = StateMachine{}
    m.state_table = make([]State, max)
    return m
}

func main() { 
    var m Monster = Monster{}
    fsm := NewStateManager(state_sizeof)
    fsm.Add(state_init, &StateInit{})
    fsm.Add(state_battle, &StateBattle{})
    fsm.SetState(state_init, &m);
    fsm.Update(&m)
    fsm.SetState(state_battle, &m);
    fsm.Update(&m)
}
