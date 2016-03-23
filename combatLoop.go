package main

import "fmt"

type combatAction struct {
	name          string
	successChance int
}

var combatOptions = map[string]*combatAction{
	"Attack!": {name: "Attack!", successChance: 50},
	"Block":   {name: "Block", successChance: 50},
	"Retreat": {name: "Retreat", successChance: 80},
}

func attack(attacker organism, h hittable, weapon *weapon) {
	var attackDamage = genRandom(weapon.baseDamage-5, weapon.baseDamage+5)
	h.takeDamage(attackDamage)
	if attacker.isPlayer == false {
		fmt.Printf("The %s hits you for %d damage\n", attacker.name, attackDamage)
	} else {
		fmt.Printf("You strike the enemy for %d damage.\n", attackDamage)
	}
	insult := genRandom(0, 99)
	if weapon.bonus.chance > insult {
		h.takeDamage(weapon.bonus.bonusDamage)
		h.updateStatus(weapon.bonus.statusEffect)
		if attacker.isPlayer == false {
			fmt.Printf("The enemy landed a %s! You took %d extra damage.\n", weapon.bonus.name, weapon.bonus.bonusDamage)
			if weapon.bonus.statusEffect.name != "No Effect" {
				fmt.Printf("Due to the %s, you are now %s!\n", weapon.bonus.name, weapon.bonus.statusEffect.name)
			}
		} else {
			fmt.Printf("You landed a %s! The enemy took %d extra damage.\n", weapon.bonus.name, weapon.bonus.bonusDamage)
			if weapon.bonus.statusEffect.name != "No Effect" {
				fmt.Printf("Due to your %s, the enemy is now %s!\n", weapon.bonus.name, weapon.bonus.statusEffect.name)
			}
		}
	}
	weapon.integrity -= genRandom(0, 5)
}

func playerTurn(player *character, mob *npc, runaway bool) bool {
	fmt.Printf("You have %d Health. Your opponent has %d.\nWhat would you like to do?\n", player.hitpoints, mob.hitpoints)
	fmt.Print("1: Attack | 2: Run away!\n\n")
	playerChoice := 0
	fmt.Scan(&playerChoice)
	if playerChoice == 1 {
		attack(player.organism, mob, player.weapon1)
	} else {
		if genRandom(5, 40) > mob.speed {
			runaway = true
		}
	}
	return runaway
}

func mobTurn(player *character, mob *npc) {
	if genRandom(1, 100) >= 30 {
		attack(mob.organism, player, mob.weapon1)
	} else {
		attack(mob.organism, player, mob.weapon2)
	}
}

func combatCycle(player *character, mob *npc, runaway bool) bool {
	fmt.Printf("You have encountered a %s!\n\n", mob.name)
	if player.size > mob.size {
		fmt.Print("You have the advantage!\n\n")
		playerTurn(player, mob, runaway)
	} else {
		fmt.Printf("The %s has the advantage!\n\n", mob.name)
		mobTurn(player, mob)
	}
	for mob.hitpoints > 0 && runaway == false && player.hitpoints > 0 {
		runaway = playerTurn(player, mob, runaway)
		mobTurn(player, mob)
	}
	if runaway == false {
		if mob.hitpoints <= 0 && player.hitpoints >= 1 {
			fmt.Printf("You defeated the %s! You gain %d experience points.\n\n", mob.name, mob.size)
			player.experience += mob.size
		}
		if player.hitpoints <= 0 {
			fmt.Printf("You were defeated by the %s!\n\n", mob.name)
		}
	} else {
		fmt.Printf("You successfully ran from the %s!\n\n", mob.name)
	}
	return runaway
}
