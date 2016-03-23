package main

import "fmt"

// Game holds initial state, basis for loop
type Game struct {
	Welcome         string
	CurrentLocation string
	playerCharacter *character
}

// CurrentLocation is... the player's current location!
var CurrentLocation *Location

// Play defines the game loop
func (g *Game) Play() {
	CurrentLocation = locationMap["Enter the clearing"]
	priorLocation := locationMap["Enter the clearing"]
	fmt.Println(g.Welcome)
	for {
		runaway := false
		fmt.Println(CurrentLocation.description) //Where are you?
		if CurrentLocation.encounterChance > genRandom(0, 99) {
			runaway = combatCycle(g.playerCharacter, encounterMob(CurrentLocation.locationEncounter), false)
		}
		if g.playerCharacter.hitpoints <= 0 { //Are you dead?
			fmt.Println("You are dead, game over!")
			return
		}
		if runaway == true {
			CurrentLocation = priorLocation
			fmt.Print("You have returned to your previous location.\n")
			fmt.Println(CurrentLocation.description)
		}
		if g.playerCharacter.hitpoints > g.playerCharacter.maxHitpoints {
			g.playerCharacter.hitpoints = g.playerCharacter.maxHitpoints
		}
		fmt.Println("Armament: ", g.playerCharacter.weapon1.name, "and", g.playerCharacter.shield.name)
		fmt.Println("Inventory Slots: ", g.playerCharacter.inventoryOne.name, g.playerCharacter.inventoryTwo.name, g.playerCharacter.inventoryThree.name, g.playerCharacter.inventoryFour.name, g.playerCharacter.inventoryFive.name, g.playerCharacter.inventorySix.name)
		fmt.Printf("Health: %d\n\n", g.playerCharacter.hitpoints) //Print health information
		fmt.Println("What would you like to do?")                 //Where can you go from here?
		for index, loc := range CurrentLocation.transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(CurrentLocation.transitions) { //What would you like to do?
			fmt.Println("Choose an action: ")
			fmt.Scan(&i)
		}
		newLoc := i - 1
		priorLocation = CurrentLocation
		CurrentLocation = locationMap[CurrentLocation.transitions[newLoc]] //Go to new location based on input
	}
}
