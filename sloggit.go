package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var agil, char, dext, inte, stre, wisd int
	rSource := rand.NewSource(time.Now().UnixNano())
	rRand := rand.New(rSource)
	accept := "N"
	for accept == "N" {
		agil = rRand.Intn(20)
		char = rRand.Intn(20)
		dext = rRand.Intn(20)
		inte = rRand.Intn(20)
		stre = rRand.Intn(20)
		wisd = rRand.Intn(20)

		fmt.Println("Agility: ", agil, "Charisma: ", char, "Dexterity: ", dext, "Intelligence: ", inte, "Strength: ", stre, "Wisdom: ", wisd)
		fmt.Println("Is this OK? (Y/N)")
		fmt.Scan(&accept) //TODO: Handle incorrect input here.
	}
	var mainCharacter = character{ //TODO: random assignment of traits, with choices for player
		organism: organism{
			name:         "Steve", //TODO: choose own name
			description:  "This is you!",
			age:          16,
			hitpoints:    100,
			maxHitpoints: 100,
			traitList: traitList{
				agility:      agil,
				charisma:     char,
				dexterity:    dext,
				intelligence: inte,
				strength:     stre,
				wisdom:       wisd,
			},
			inventoryOne:   equipmentItems["Empty"],
			inventoryTwo:   equipmentItems["Empty"],
			inventoryThree: equipmentItems["Empty"],
			inventoryFour:  equipmentItems["Empty"],
			inventoryFive:  equipmentItems["Empty"],
			inventorySix:   equipmentItems["Empty"],
		},
		weapon: weaponry["Rusty Sword"],
		shield: defences["Battered Shield"],
	}

	var introduction = "You wake up in a clearing in the forest, resting on a thick carpet of pine needles. \nYou stand up and look about you, your hand instinctively reaching \nfor your Rusty Sword and Battered Shield. There is a chill in the air; \nfrom that and the trees you deduce that you must be many hundreds of miles \nfurther North than where you laid down to sleep. Suspecting foul play, \nyou check that you are undamaged and set off through the trees...\n\n"

	g := &Game{Welcome: introduction, playerCharacter: mainCharacter}
	g.Play()
}
