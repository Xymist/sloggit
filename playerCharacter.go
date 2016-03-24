package main

import (
	"fmt"
	"math"
)

type character struct { //Player characters. Currently singular, but may not be later.
	organism
	weapon1 *weapon
	*shield
	size       int
	inventory  []*equipment
	experience int
}

func generateCharacter() *character {
	myName := "Steve"
	fmt.Println("Please enter your name:")
	fmt.Scan(&myName)
	fmt.Println("Your name is ", myName)

	var agil, char, dext, inte, stre, wisd float64
	accept := "N"
	for accept == "N" || accept == "n" {
		agil = float64(genRandom(3, 20))
		char = float64(genRandom(3, 20))
		dext = float64(genRandom(3, 20))
		inte = float64(genRandom(3, 20))
		stre = float64(genRandom(3, 20))
		wisd = float64(genRandom(3, 20))

		fmt.Println("Agility: ", agil, "Charisma: ", char, "Dexterity: ", dext, "Intelligence: ", inte, "Strength: ", stre, "Wisdom: ", wisd)
		fmt.Println("Is this OK? (y/N)")
		fmt.Scan(&accept)
	}

	var generatedCharacter = &character{
		organism: organism{
			name:         myName,
			description:  "This is you!",
			isPlayer:     true,
			age:          16,
			hitpoints:    math.Floor(100.0 + (stre * 1.5)),
			maxHitpoints: math.Floor(100.0 + (stre * 1.5)),
			traitList: traitList{
				agility:      agil,
				charisma:     char,
				dexterity:    dext,
				intelligence: inte,
				strength:     stre,
				wisdom:       wisd,
			},
		},
		weapon1:    weaponry["Rusty Sword"],
		shield:     defences["Battered Shield"],
		inventory:  []*equipment{},
		size:       1,
		experience: 0,
	}
	return generatedCharacter
}
