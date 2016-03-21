package main

// Location is each place the player could exist
type Location struct {
	Description string
	Transitions []string
	Events      []string
}

var locationMap = map[string]*Location{
	"startClearing": {"You are standing in a quiet clearing. The impression of a sleeping body has been left in the soft forest floor.", []string{"Go North", "Go South", "Go East", "Go West"}, []string{"alienAttack"}},
	"Ready Room":    {"The Captain's ready room.", []string{"Bridge"}, []string{}},
	"Turbo Lift":    {"A Turbo Lift that takes you anywhere in the ship.", []string{"Bridge", "Lounge", "Engineering", "Dungeon"}, []string{"android"}},
	"Engineering":   {"You are in engineering where you see the star drive", []string{"Turbo Lift"}, []string{"alienAttack"}},
	"Lounge":        {"You are in the lounge, you feel very relaxed", []string{"Turbo Lift"}, []string{"relaxing"}},
	"Dungeon":       {"This doesn't belong here...", []string{"Turbo Lift"}, []string{"alienAttack"}},
}
