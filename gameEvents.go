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
	"monsterAttack":  {Chance: 20, Description: "A vicious gremlin slashes at you from a nearby shadow.\n", Health: -10, Evt: ""},
	"bearAttack":     {Chance: 80, Description: "You have disturbed a sleeping bear! It lunges towards you, swinging a heavy paw towards your head.\n", Health: -60, Evt: ""},
	"cliffSplatter":  {Chance: 100, Description: "You fall precipitously, your equipment tumbling away from you. The sharp rocks at the bottom rush up to meet you...\n", Health: -1000, Evt: ""},
	"fishJump":       {Chance: 25, Description: "You hear a splash as a silver fish snaps an insect from the surface of the water. There could be food here.\n", Health: 0, Evt: ""},
	"guardianAttack": {Chance: 95, Description: "There is a series of loud crunching sounds from the ceiling, and a great portion of it collapses, throwing stone fragments and dust across the chamber. As the dust clears, you see a vast stone golem crouching in the rubble. With a roar, it charges you.\n", Health: -95, Evt: ""},
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
