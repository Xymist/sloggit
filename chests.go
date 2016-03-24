package main

import "fmt"

// 'Chests' contains all objects which contain equipment, whether it's a jar, a bookcase, a desk or actually a chest.

type chest struct {
	physObject
	contents []string
}

var chests = map[string]*chest{
	"Cave Chest": {physObject: physObject{name: "Cave Chest", integrity: 100}, contents: []string{"Wooden Shield", "Torch", "Climbing Rope", "Tinderbox"}},
}

func openChest(c *chest, player *character) {
	fmt.Println("You see something inside:")
	fmt.Print("\t0 - Leave everything\n")
	for index, element := range c.contents {
		fmt.Printf("\t%d - Take %s\n", index+1, element)
	}
	choice := -1
	for choice < 0 || choice > len(c.contents) { //Loop until they pick something within bounds
		fmt.Println("Choose action: ")
		fmt.Scan(&choice)
	}
	if choice == 0 {
		fmt.Print("There was nothing interesting here.\n")
		return
	}
	choice = choice - 1
	newContents := []string{}
	if equipmentItems[c.contents[choice]] != nil {
		loot := equipmentItems[c.contents[choice]]
		player.inventory = append(player.inventory, loot)
		fmt.Printf("You receive a %s! You put the %s in your backpack.\n", loot.name, loot.name)
	}
	if defences[c.contents[choice]] != nil {
		player.shield = defences[c.contents[choice]]
		fmt.Printf("Shield upgraded to %s!\n", player.shield.name)
	}
	if weaponry[c.contents[choice]] != nil {
		player.weapon1 = weaponry[c.contents[choice]]
		fmt.Printf("Weapon upgraded to %s!\n", player.weapon1.name)
	}
	for _, item := range c.contents {
		if item != c.contents[choice] {
			newContents = append(newContents, item)
		}
	}
	c.contents = newContents
}
