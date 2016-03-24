package main

// Location is each place the player could exist
type Location struct {
	description       string
	transitions       []string
	locationEncounter string
	encounterChance   int
}

var locationMap = map[string]*Location{
	"Enter the clearing":             {description: clearingText, transitions: []string{"Go North", "Go South", "Go East", "Go West"}, locationEncounter: "Monster Attack", encounterChance: 0},
	"Go North":                       {description: startNorthText, transitions: []string{"Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go South":                       {description: startSouthText, transitions: []string{"Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go East":                        {description: startEastText, transitions: []string{"Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Go West":                        {description: startWestText, transitions: []string{"Follow the river upstream", "Follow the river downstream", "Enter the clearing"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Follow the river upstream":      {description: upRiverText, transitions: []string{"Duck into the cave", "Look for a good tree to climb", "Head further North"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Follow the river downstream":    {description: downRiverText, transitions: []string{"Leap from the cliff"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Leap from the cliff":            {description: cliffJumpText, transitions: []string{}, locationEncounter: "Cliff Dive", encounterChance: 100},
	"Duck into the cave":             {description: caveEntranceText, transitions: []string{"Open Cave Chest", "Venture further into the cave", "Leave cave"}, locationEncounter: "Grizzly Attack", encounterChance: 40},
	"Open Cave Chest":                {description: caveChestText, transitions: []string{"Open Cave Chest", "Venture further into the cave", "Leave cave"}, locationEncounter: "Cave Chest", encounterChance: 100},
	"Venture further into the cave":  {description: caveDepthsText, transitions: []string{"Back to cave entrance"}, locationEncounter: "Monster Attack", encounterChance: 50},
	"Back to cave entrance":          {description: caveEntranceText, transitions: []string{"Open Cave Chest", "Venture further into the cave", "Leave cave"}, locationEncounter: "Monster Attack", encounterChance: 20},
	"Leave cave":                     {description: caveExitText, transitions: []string{"Duck into the cave", "Look for a good tree to climb", "Head further North"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Look for a good tree to climb":  {description: treeClimbText, transitions: []string{}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Head further North":             {description: farNorthText, transitions: []string{}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Enter the dungeon":              {description: dungeonEntranceText, transitions: []string{"Fling yourself into the portal"}, locationEncounter: "Monster Attack", encounterChance: 30},
	"Fling yourself into the portal": {description: portalFallText, transitions: []string{}, locationEncounter: "Monster Attack", encounterChance: 30},
}
