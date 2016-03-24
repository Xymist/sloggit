package main

type npc struct { //Includes things like motile trees; any agent that is not a player character.
	organism
	hostile bool
	weapon1 *weapon
	weapon2 *weapon
	*shield
	size  int
	speed int
}

func (n *npc) takeDamage(damage int) {
	n.hitpoints = n.hitpoints - damage
}

func (n *npc) updateStatus(statusEffect *statusEffect) {
	n.statusEffect = statusEffect
}

var hostileMobs = map[string]*npc{
	"Small Kobold": {organism: organism{
		name:        "Small Kobold",
		description: "Tiny creature, huge problem if they get into knife range",
		hitpoints:   20,
	},
		hostile: true,
		weapon1: weaponry["Rusty Dagger"],
		weapon2: weaponry["Rusty Dagger"],
		size:    1,
		speed:   6,
	},
	"Kobold": {organism: organism{
		name:        "Kobold",
		description: "This adult kobold still has trouble lifting his stolen sword",
		hitpoints:   35,
	},
		hostile: true,
		weapon1: weaponry["Rusty Sword"],
		weapon2: weaponry["Rusty Sword"],
		size:    2,
		speed:   8,
	},
	"Grizzly Bear": {organism: organism{
		name:        "Grizzly Bear",
		description: "This creature would be peaceful, if you had stayed away",
		hitpoints:   60,
	},
		hostile: true,
		weapon1: weaponry["Bear Paw"],
		weapon2: weaponry["Bear Jaws"],
		size:    5,
		speed:   32,
	},
	"Craggy Cliff": {organism: organism{
		name:        "Craggy Cliff",
		description: "Gravity is not your friend.",
		hitpoints:   75,
	},
		hostile: true,
		weapon1: weaponry["Spiked Rocks"],
		weapon2: weaponry["Spiked Rocks"],
		size:    5,
		speed:   0,
	},
}
