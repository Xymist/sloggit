package main

type traitList struct {
	charisma, intelligence, wisdom, strength, dexterity, agility int
}

type organism struct {
	name, description string
	isPlayer          bool
	age               int
	hitpoints         int
	maxHitpoints      int
	traitList
	statusEffect *statusEffect
}

func (o *organism) takeDamage(damage int) {
	o.hitpoints -= damage
}
