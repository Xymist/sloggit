package main

// Location is each place the player could exist
type Location struct {
	description       string
	chests            []string
	transitions       []string
	locationEncounter string
	encounterChance   int
}

var locationMap = map[string]*Location{
	"Enter the clearing":             {description: clearingText, chests: []string{}, transitions: []string{"Go North", "Go South", "Go East", "Go West"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go North":                       {description: startNorthText, chests: []string{}, transitions: []string{"Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go South":                       {description: startSouthText, chests: []string{}, transitions: []string{"Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go East":                        {description: startEastText, chests: []string{}, transitions: []string{"Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go West":                        {description: startWestText, chests: []string{}, transitions: []string{"Follow the river upstream", "Follow the river downstream", "Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Follow the river upstream":      {description: upRiverText, chests: []string{}, transitions: []string{"Duck into the cave", "Look for a good tree to climb", "Head further North"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Follow the river downstream":    {description: downRiverText, chests: []string{}, transitions: []string{"Leap from the cliff"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Leap from the cliff":            {description: cliffJumpText, chests: []string{}, transitions: []string{}, locationEncounter: "Cliff Dive", encounterChance: 100},
	"Duck into the cave":             {description: caveEntranceText, chests: []string{"Cave Chest"}, transitions: []string{}, locationEncounter: "Grizzly Attack", encounterChance: 80},
	"Look for a good tree to climb":  {description: treeClimbText, chests: []string{}, transitions: []string{}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Head further North":             {description: farNorthText, chests: []string{}, transitions: []string{}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Enter the dungeon":              {description: dungeonEntranceText, chests: []string{}, transitions: []string{"Fling yourself into the portal"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Fling yourself into the portal": {description: portalFallText, chests: []string{}, transitions: []string{}, locationEncounter: "Monster Attack", encounterChance: 30},
}
