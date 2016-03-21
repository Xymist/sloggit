package main

type hittable interface {
	takeDamage(int)
}

type entity interface {
	inspect()
}

func attack(h hittable, damage int, weapon weapon) {
	h.takeDamage(damage)
	weapon.integrity--
}

func examine(e entity) {
	e.inspect()
}
