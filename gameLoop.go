package main

import (
	"fmt"
	"math"
)

// Game holds initial state, basis for loop
type game struct {
	Welcome         string
	currentLocation string
	playerCharacter *character
}

// currentLocation is... the player's current location!
var currentLocation *Location

// Play defines the game loop
func (g *game) Play() {
	currentLocation = locationMap["Enter the clearing"]
	priorLocation := locationMap["Enter the clearing"]
	player := g.playerCharacter
	fmt.Println(g.Welcome)
	for {
		runaway := false
		fmt.Println(currentLocation.description) //Where are you?
		priorLevel := player.size
		player.size = int(math.Floor(math.Sqrt(float64(player.experience))))
		fmt.Printf("Armament: %s and %s\n", player.weapon1.name, player.shield.name)
		fmt.Println("Inventory Slots: ")
		for index, invItem := range player.inventory {
			fmt.Printf("%d: %s, ", index+1, invItem.name)
		}
		fmt.Printf("\nHealth: %.f\n", player.hitpoints) //Print health information
		if player.size > priorLevel {
			fmt.Printf("You levelled up! New level: %d\n\n", player.size)
		} else {
			fmt.Printf("Level: %d\n\n", player.size)
		}
		if currentLocation.encounterChance > genRandom(0, 99) {
			if encounterMob(currentLocation.locationEncounter) != nil {
				runaway = combatCycle(player, encounterMob(currentLocation.locationEncounter), false)
			} else {
				openChest(chests[events[currentLocation.locationEncounter].name], player)
			}
		}
		if player.hitpoints <= 0 { //Are you dead?
			fmt.Println("You are dead, game over!")
			return
		}
		if runaway == true {
			currentLocation = priorLocation
			fmt.Print("You have returned to your previous location.\n")
			fmt.Println(currentLocation.description)
		}
		if player.hitpoints > player.maxHitpoints {
			player.hitpoints = player.maxHitpoints
		}
		fmt.Println("What would you like to do?") //Where can you go from here?
		for index, loc := range currentLocation.transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(currentLocation.transitions) { //What would you like to do?
			fmt.Println("Choose an action: ")
			fmt.Scan(&i)
		}
		callClear()
		newLoc := i - 1
		priorLocation = currentLocation
		currentLocation = locationMap[currentLocation.transitions[newLoc]] //Go to new location based on input
	}
}
