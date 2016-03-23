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
}

func (o *organism) takeDamage(damage int) {
	o.hitpoints -= damage
}
