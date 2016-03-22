package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateCharacter() character {
	myName := "Steve"
	fmt.Println("Please enter your name:")
	fmt.Scan(&myName)
	fmt.Println("Your name is ", myName)

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

	var generatedCharacter = character{
		organism: organism{
			name:         myName,
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
	return generatedCharacter
}
