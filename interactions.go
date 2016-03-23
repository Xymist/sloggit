package main

type hittable interface {
	takeDamage(int)
}

type entity interface {
	inspect()
}

func attack(h hittable, weapon weapon) {
	h.takeDamage(weapon.baseDamage)
	weapon.integrity--
}

func examine(e entity) {
	e.inspect()
}
