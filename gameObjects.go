package main

import (
	"fmt"
)

type physObject struct {
	name      string
	integrity int
}

func (p *physObject) takeDamage(damage int) {
	p.integrity -= damage
}

type weapon struct {
	physObject
	baseDamage int
	bonuses    []string
}

func (w *weapon) inspect() {
	fmt.Println(w.name, "Integrity: ", w.integrity, "Damage: ", w.baseDamage)
}

type shield struct {
	physObject
	bonuses []string
}

func (s *shield) inspect() {
	fmt.Println(s.name, "Integrity: ", s.integrity)
}

type chest struct {
	physObject
	contents []string
}
