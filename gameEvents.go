package main

// Event may happen to a player
type event struct {
	name string
	npcs []string
}

func encounterMob(s string) *npc {
	evt := events[s]
	chooseNPC := genRandom(1, len(evt.npcs))
	return hostileMobs[evt.npcs[chooseNPC]]
}

var events = map[string]*event{
	"Monster Attack": {name: "Monster Attack", npcs: []string{"Small Kobold", "Kobold"}},
	"Grizzly Attack": {name: "Grizzly Attack", npcs: []string{"Grizzly Bear"}},
}
