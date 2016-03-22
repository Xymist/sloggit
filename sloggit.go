package main

func main() {

	var mainCharacter = character{ //TODO: random assignment of traits, with choices for player
		organism: organism{
			name:        "Steve", //TODO: choose own name
			description: "This is you!",
			age:         16,
			hitpoints:   100,
			traitList: traitList{
				agility:      5,
				charisma:     5,
				dexterity:    5,
				intelligence: 5,
				strength:     5,
				wisdom:       5,
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

	g := &Game{Health: 100, Welcome: introduction, maxHealth: 100, playerCharacter: mainCharacter}
	g.Play()
}
