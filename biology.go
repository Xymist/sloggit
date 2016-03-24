package main

type traitList struct {
	charisma, intelligence, wisdom, strength, dexterity, agility float64
}

type organism struct {
	name, description string
	isPlayer          bool
	age               int
	hitpoints         float64
	maxHitpoints      float64
	traitList
	statusEffect *statusEffect
}

func (o *organism) takeDamage(damage int) {
	o.hitpoints -= float64(damage)
}

func (o *organism) updateStatus(statusEffect *statusEffect) {
	o.statusEffect = statusEffect
}
