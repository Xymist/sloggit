package main

type event struct {
	name  string
	npcs  []string
	chest string
}

func encounterMob(s string) *npc {
	evt := events[s]
	chooseNPC := 0
	if len(evt.npcs) > 0 {
		if len(evt.npcs) > 1 {
			chooseNPC = genRandom(1, len(evt.npcs)) - 1
		}
		return hostileMobs[evt.npcs[chooseNPC]]
	}
	return nil
}

var events = map[string]*event{
	"Monster Attack":         {name: "Monster Attack", npcs: []string{"Small Kobold", "Kobold"}},
	"Grizzly Attack":         {name: "Grizzly Attack", npcs: []string{"Grizzly Bear"}},
	"Cliff Dive":             {name: "Cliff Dive", npcs: []string{"Craggy Cliff"}},
	"Cave Chest":             {name: "Cave Chest", npcs: []string{}, chest: "Cave Chest"},
	"Tree Hollow":            {name: "Tree Hollow", npcs: []string{}, chest: "Tree Hollow"},
	"Treecreature Encounter": {name: "Treecreature Encounter", npcs: []string{"Vicious Squirrel"}},
}
