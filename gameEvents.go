package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Event may happen to a player
type Event struct {
	Type        string
	Chance      int
	Description string
	Health      int
	Evt         string
}

var evts = map[string]*Event{
	"monsterAttack": {Chance: 20, Description: "A vicious gremlin slashes at you from a nearby shadow.", Health: -50, Evt: ""},
}

// ProcessEvent allows the game to notice what happened
func (e *Event) ProcessEvent() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if e.Chance >= r1.Intn(100) {
		if e.Type == "Combat" {
			fmt.Println("Combat Event")
		}
		fmt.Printf("\t%s\n", e.Description)
		if e.Evt != "" {
			e.Health = e.Health + evts[e.Evt].ProcessEvent()
		}
		return e.Health
	}
	return 0
}
