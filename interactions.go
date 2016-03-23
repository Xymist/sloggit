package main

type hittable interface {
	takeDamage(int)
	updateStatus(*statusEffect)
}

type entity interface {
	inspect()
}

func attack(h hittable, weapon weapon) {
	h.takeDamage(weapon.baseDamage)
	insult := genRandom(0, 99)
	if weapon.bonus.chance > insult {
		h.takeDamage(weapon.bonus.bonusDamage)
		h.updateStatus(weapon.bonus.statusEffect)
	}
	weapon.integrity--
}

func examine(e entity) {
	e.inspect()
}
