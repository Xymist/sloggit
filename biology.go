package main

type traitList struct {
	charisma, intelligence, wisdom, strength, dexterity, agility int
}

type organism struct {
	name, description string
	age               int
	hitpoints         int
	traitList
	inventoryOne, inventoryTwo, inventoryThree, inventoryFour, inventoryFive, inventorySix *equipment
}

func (o *organism) takeDamage(damage int) {
	o.hitpoints -= damage
}

type character struct {
	organism
	*weapon
	*shield
}

type npc struct {
	organism
	hostile bool
	*weapon
	*shield
}
