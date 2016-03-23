package main

// Location is each place the player could exist
type Location struct {
	Description string
	Chests      []string
	Transitions []string
	Events      []string
}

var locationMap = map[string]*Location{
	"Enter the clearing":             {clearingText, []string{}, []string{"Go North", "Go South", "Go East", "Go West"}, []string{"monsterAttack"}},
	"Go North":                       {startNorthText, []string{}, []string{"Enter the clearing"}, []string{}},
	"Go South":                       {startSouthText, []string{}, []string{"Enter the clearing"}, []string{}},
	"Go East":                        {startEastText, []string{}, []string{"Enter the clearing"}, []string{}},
	"Go West":                        {startWestText, []string{}, []string{"Follow the river upstream", "Follow the river downstream", "Enter the clearing"}, []string{"fishJump"}},
	"Follow the river upstream":      {upRiverText, []string{}, []string{"Duck into the cave", "Look for a good tree to climb", "Head further North"}, []string{}},
	"Follow the river downstream":    {downRiverText, []string{}, []string{"Leap from the cliff"}, []string{}},
	"Leap from the cliff":            {cliffJumpText, []string{}, []string{}, []string{"cliffSplatter"}},
	"Duck into the cave":             {caveEntranceText, []string{"Cave Chest"}, []string{}, []string{"bearAttack"}},
	"Look for a good tree to climb":  {treeClimbText, []string{}, []string{}, []string{}},
	"Head further North":             {farNorthText, []string{}, []string{}, []string{}},
	"Enter the dungeon":              {dungeonEntranceText, []string{}, []string{"Fling yourself into the portal"}, []string{"guardianAttack"}},
	"Fling yourself into the portal": {portalFallText, []string{}, []string{}, []string{"homecoming"}},
}
