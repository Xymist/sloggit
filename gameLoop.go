package main

import "fmt"

// Game holds initial state, basis for loop
type Game struct {
	Welcome         string
	CurrentLocation string
	playerCharacter character
}

// CurrentLocation is... the player's current location!
var CurrentLocation *Location

// Play defines the game loop
func (g *Game) Play() {
	CurrentLocation = locationMap["Enter the clearing"]
	fmt.Println(g.Welcome)
	for {
		fmt.Println(CurrentLocation.Description) //Where are you?
		g.ProcessEvents(CurrentLocation.Events)  //Did anything happen here?
		if g.playerCharacter.hitpoints < 0 {     //Are you dead?
			fmt.Println("You are dead, game over!")
			return
		}
		if g.playerCharacter.hitpoints > g.playerCharacter.maxHitpoints {
			g.playerCharacter.hitpoints = g.playerCharacter.maxHitpoints
		}
		fmt.Println("Armament: ", g.playerCharacter.weapon.name, "and", g.playerCharacter.shield.name)
		fmt.Println("Inventory Slots: ", g.playerCharacter.inventoryOne.name, g.playerCharacter.inventoryTwo.name, g.playerCharacter.inventoryThree.name, g.playerCharacter.inventoryFour.name, g.playerCharacter.inventoryFive.name, g.playerCharacter.inventorySix.name)
		fmt.Printf("Health: %d\n\n", g.playerCharacter.hitpoints) //Print health information
		fmt.Println("What would you like to do?")                 //Where can you go from here?
		for index, loc := range CurrentLocation.Transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(CurrentLocation.Transitions) { //What would you like to do?
			fmt.Println("Choose an action: ")
			fmt.Scan(&i)
		}
		newLoc := i - 1
		CurrentLocation = locationMap[CurrentLocation.Transitions[newLoc]] //Go to new location based on input

	}
}

// ProcessEvents is the command to process events occurring this loop
func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.playerCharacter.hitpoints += evts[evtName].ProcessEvent()
	}
}
