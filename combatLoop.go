package main

import (
	"fmt"

	"github.com/ttacon/chalk"
)

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
	fmt.Printf("%sYou have encountered %s!%s", chalk.Red, n.name, chalk.Reset)
	fmt.Println("What would you like to do?")
	runaway := false
	for n.hitpoints >= 0 || runaway == true { //Loop combat until the opponent is dead or the player escapes

	}

}
