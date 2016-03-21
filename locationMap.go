package main

// Location is each place the player could exist
type Location struct {
	Description string
	Transitions []string
	Events      []string
}

var locationMap = map[string]*Location{
	"startClearing": {"You are standing in a quiet clearing. The impression of a sleeping body has been left in the soft forest floor.", []string{"Go North", "Go South", "Go East", "Go West"}, []string{"monsterAttack"}},
	"Dungeon":       {"This doesn't belong here...", []string{"Turbo Lift"}, []string{"monsterAttack"}},
}
