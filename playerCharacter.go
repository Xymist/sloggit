package main

import "fmt"

type character struct { //Player characters. Currently singular, but may not be later.
	organism
	weapon1 *weapon
	*shield
	size       int
	inventory  []*equipment
	experience int
}

func (c *character) takeDamage(damage int) {
	c.hitpoints = c.hitpoints - damage
}

func (c *character) updateStatus(statusEffect *statusEffect) {
	c.statusEffect = statusEffect
}

func generateCharacter() *character {
	myName := "Steve"
	fmt.Println("Please enter your name:")
	fmt.Scan(&myName)
	fmt.Println("Your name is ", myName)

	var agil, char, dext, inte, stre, wisd int
	accept := "N"
	for accept == "N" || accept == "n" {
		agil = genRandom(3, 20)
		char = genRandom(3, 20)
		dext = genRandom(3, 20)
		inte = genRandom(3, 20)
		stre = genRandom(3, 20)
		wisd = genRandom(3, 20)

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
		},
		weapon1:    weaponry["Rusty Sword"],
		shield:     defences["Battered Shield"],
		inventory:  []*equipment{},
		size:       1,
		experience: 0,
	}
	return generatedCharacter
}
