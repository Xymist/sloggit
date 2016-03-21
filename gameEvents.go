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
	"alienAttack":     {Chance: 20, Description: "An alien beams in front of you and shoots you with a ray gun.", Health: -50, Evt: "doctorTreatment"},
	"doctorTreatment": {Chance: 10, Description: "The doctor rushes in and inject you with a health boost.", Health: +30, Evt: ""},
	"android":         {Chance: 50, Description: "Data is in the turbo lift and says hi to you", Health: 0, Evt: ""},
	"relaxing":        {Chance: 100, Description: "In the lounge you are so relaxed that your health improves.", Health: +10, Evt: ""},
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
