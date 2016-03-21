package main

import "fmt"

// Game holds initial state, basis for loop
type Game struct {
	Welcome         string
	Health          int
	maxHealth       int
	CurrentLocation string
	playerCharacter character
}

// CurrentLocation is... the player's current location!
var CurrentLocation *Location

// Play is the game loop
func (g *Game) Play() {
	CurrentLocation = locationMap["startClearing"]
	fmt.Println(g.Welcome)
	for {
		fmt.Println(CurrentLocation.Description) //Where are you?
		g.ProcessEvents(CurrentLocation.Events)  //Did anything happen here?
		if g.Health < 0 {                        //Are you dead?
			fmt.Println("You are dead, game over!!!")
			return
		}
		if g.Health > g.maxHealth {
			g.Health = g.maxHealth
		}
		fmt.Println(g.playerCharacter.weapon.name)
		fmt.Printf("Health: %d\n", g.Health)      //Print health information
		fmt.Println("What would you like to do?") //Where can you go from here?
		for index, loc := range CurrentLocation.Transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(CurrentLocation.Transitions) { //What would you like to do?
			fmt.Printf("%s%d%s\n", "Choose a direction. (0 - to quit), [1...", len(CurrentLocation.Transitions), "]: ")
			fmt.Scan(&i)
		}
		newLoc := i - 1
		CurrentLocation = locationMap[CurrentLocation.Transitions[newLoc]] //Go to new location based on input

	}
}

// ProcessEvents is the command to process events occurring this loop
func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Health += evts[evtName].ProcessEvent()
	}
}
