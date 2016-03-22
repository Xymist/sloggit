package main

type traitList struct {
	charisma, intelligence, wisdom, strength, dexterity, agility int
}

type organism struct {
	name, description string
	age               int
	hitpoints         int
	maxHitpoints      int
	traitList
	inventoryOne, inventoryTwo, inventoryThree, inventoryFour, inventoryFive, inventorySix *equipment
}

func (o *organism) takeDamage(damage int) {
	o.hitpoints -= damage
}

type character struct { //Player characters. Currently singular, but may not be later.
	organism
	*weapon
	*shield
}

type npc struct { //Includes things like motile trees; any agent that is not a player character.
	organism
	hostile bool
	*weapon
	*shield
}
