package main

import "fmt"

//Weapons are anything used to cause damage. Swords, wands, firearms etc

type weapon struct {
	physObject
	baseDamage int
	bonus      *damageBonus
}

type damageBonus struct {
	name         string
	bonusDamage  int
	chance       int
	statusEffect *statusEffect
}

type statusEffect struct {
	name             string
	duration         int
	maxHealthPenalty int
}

func (w *weapon) inspect() {
	fmt.Println(w.name, "Integrity: ", w.integrity, "Damage: ", w.baseDamage, "Bonus Effects: ", w.bonus)
}

var weaponry = map[string]*weapon{
	"Rusty Sword":  {physObject: physObject{name: "Rusty Sword", integrity: 20}, baseDamage: 20, bonus: damageBonuses["No Bonus"]},
	"Rusty Dagger": {physObject: physObject{name: "Rusty Dagger", integrity: 15}, baseDamage: 12, bonus: damageBonuses["No Bonus"]},
	"Bear Paw":     {physObject: physObject{name: "Bear Paw", integrity: 100}, baseDamage: 20, bonus: damageBonuses["Vicious Blow"]},
	"Bear Jaws":    {physObject: physObject{name: "Bear Jaws", integrity: 100}, baseDamage: 40, bonus: damageBonuses["Crushing Bite"]},
}

var damageBonuses = map[string]*damageBonus{
	"No Bonus":      {name: "No Bonus", bonusDamage: 0, chance: 0, statusEffect: statusOptions["No Effect"]},
	"Vicious Blow":  {name: "Vicious Blow", bonusDamage: 10, chance: 33, statusEffect: statusOptions["No Effect"]},
	"Crushing Bite": {name: "Crushing Bite", bonusDamage: 20, chance: 18, statusEffect: statusOptions["Crippled"]},
}

var statusOptions = map[string]*statusEffect{
	"No Effect": {name: "No Effect", duration: 0, maxHealthPenalty: 0},
	"Crippled":  {name: "Crippled", duration: 50, maxHealthPenalty: 30},
}
