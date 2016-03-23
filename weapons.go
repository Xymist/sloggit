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
	"Rusty Sword":  {physObject: physObject{name: "Rusty Sword", integrity: 100}, baseDamage: 20, bonuses: []string{"No Bonus"}},
	"Rusty Dagger": {physObject: physObject{name: "Rusty Dagger", integrity: 60}, baseDamage: 12, bonuses: []string{"No Bonus"}},
	"Bear Paw":     {physObject: physObject{name: "Bear Paw", integrity: 100}, baseDamage: 20, bonuses: []string{"No Bonus"}},
	"Bear Jaws":    {physObject: physObject{name: "Bear Jaws", integrity: 100}, baseDamage: 40, bonuses: []string{"No Bonus"}},
}
