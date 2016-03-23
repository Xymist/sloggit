package main

import "fmt"

//Weapons are anything used to cause damage. Swords, wands, firearms etc

type weapon struct {
	physObject
	baseDamage int
	bonuses    []string
}

func (w *weapon) inspect() {
	fmt.Println(w.name, "Integrity: ", w.integrity, "Damage: ", w.baseDamage)
}

var weaponry = map[string]*weapon{
	"Rusty Sword": {physObject: physObject{name: "Rusty Sword", integrity: 100}, baseDamage: 3, bonuses: []string{"No Bonus"}},
}
