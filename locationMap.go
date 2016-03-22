package main

// Location is each place the player could exist
type Location struct {
	Description string
	Objects     []string
	Transitions []string
	Events      []string
}

var locationMap = map[string]*Location{
	"Enter the clearing":             {"You are standing in a quiet clearing. The impression of a sleeping body has been left in the soft forest floor.\n", []string{}, []string{"Go North", "Go South", "Go East", "Go West"}, []string{"monsterAttack"}},
	"Go North":                       {"", []string{}, []string{}, []string{}},
	"Go South":                       {"", []string{}, []string{}, []string{}},
	"Go East":                        {"", []string{}, []string{}, []string{}},
	"Go West":                        {"You walk through the trees until you find yourself facing a river. It's far too fast to ford; you could follow it and look for a crossing, but you might do better to go back to the clearing.\n", []string{}, []string{"Follow the river upstream", "Follow the river downstream", "Enter the clearing"}, []string{"fishJump"}},
	"Follow the river upstream":      {"You make your way along the riverbank, heading roughly North according to the sun. After several hours it begins to get dark, but you're yet to find a crossing. You start looking for shelter in the surrounding landscape, and shortly come across a cave. Wilderness caves are risky, but it could represent safety.\n", []string{}, []string{"Duck into the cave", "Look for a good tree to climb", "Head further North"}, []string{}},
	"Follow the river downstream":    {"This path is steep and rocky, and after some time you begin to hear a faint rumble in the distance, which becomes a muffled roar as you continue to follow the river. Presently you find yourself standing at the side of a waterfall as the river cascades over a tall cliff, the ground beneath obscured by cloud and spray.\n", []string{}, []string{"Leap from the cliff"}, []string{}},
	"Leap from the cliff":            {"That was daft, wasn't it?\n", []string{}, []string{}, []string{"cliffSplatter"}},
	"Duck into the cave":             {"It is cool and dark in here, though the floor is dry dirt. The sounds from outside are muffled. You wish for a torch, but through the gloom you spy a banded chest against the wall.", []string{"Cave Chest"}, []string{}, []string{"bearAttack"}},
	"Look for a good tree to climb":  {"", []string{}, []string{}, []string{}},
	"Head further North":             {"", []string{}, []string{}, []string{}},
	"Enter the dungeon":              {"This doesn't belong here...\n", []string{}, []string{"Fling yourself into the portal"}, []string{"guardianAttack"}},
	"Fling yourself into the portal": {"You take a deep breath, and let yourself fall towards the swirling energies beneath you. There is a flash of light, and you are knocked unconscious.\n", []string{}, []string{}, []string{"homecoming"}},
}
