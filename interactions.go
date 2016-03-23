package main

import "fmt"

type hittable interface {
	takeDamage(int)
	updateStatus(*statusEffect)
}

type entity interface {
	inspect()
}

func attack(attacker organism, h hittable, weapon weapon) {
	var attackDamage = genRandom(weapon.baseDamage-5, weapon.baseDamage+5)
	h.takeDamage(attackDamage)
	if attacker.isPlayer == false {
		fmt.Printf("The %s hits you for %d damage", attacker.name, attackDamage)
	} else {
		fmt.Printf("You strike the enemy for %d", attackDamage)
	}
	insult := genRandom(0, 99)
	if weapon.bonus.chance > insult {
		h.takeDamage(weapon.bonus.bonusDamage)
		h.updateStatus(weapon.bonus.statusEffect)
		if attacker.isPlayer == false {
			fmt.Printf("The enemy landed a %s! You took %d extra damage and are now %s!", weapon.bonus.name, weapon.bonus.bonusDamage, weapon.bonus.statusEffect)
		} else {
			fmt.Printf("You landed a %s! The enemy took %d extra damage and is now %s!", weapon.bonus.name, weapon.bonus.bonusDamage, weapon.bonus.statusEffect)
		}
	}
	weapon.integrity -= genRandom(0, 5)
}

func examine(e entity) {
	e.inspect()
}
