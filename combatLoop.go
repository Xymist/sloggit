package main

import "fmt"

type combatAction struct {
	name          string
	successChance int
}

var combatOptions = map[string]*combatAction{
	"Attack!": {name: "Attack!", successChance: 50},
	"Block":   {name: "Block", successChance: 50},
	"Retreat": {name: "Retreat", successChance: 80},
}

func combatLoop(n npc) {
	fmt.Printf("You have encountered %s!", n.name)
	fmt.Println("What would you like to do?")
	runaway := false
	for n.hitpoints >= 0 || runaway == true { //Loop combat until the opponent is dead or the player escapes

	}

}
