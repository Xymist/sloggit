package main

type npc struct { //Includes things like motile trees; any agent that is not a player character.
	organism
	hostile bool
	weapon1 *weapon
	weapon2 *weapon
	*shield
	size int
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
	},
		hostile: true,
		weapon1: weaponry["Rusty Dagger"],
		size:    1,
	},
	"Kobold": {organism: organism{
		name:        "Kobold",
		description: "This adult kobold still has trouble lifting his stolen sword",
	},
		hostile: true,
		weapon1: weaponry["Rusty Sword"],
		size:    2,
	},
	"Grizzly Bear": {organism: organism{
		name:        "Grizzly Bear",
		description: "This creature would be peaceful, if you had stayed away",
	},
		hostile: true,
		weapon1: weaponry["Bear Paw"],
		weapon2: weaponry["Bear Jaws"],
		size:    5,
	},
}
