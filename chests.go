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
	for index, element := range c.contents {
		fmt.Printf("\t%d - %s\n", index+1, element)
	}
	choice := 0
	for choice < 1 || choice > len(c.contents) { //Loop until they pick something within bounds
		fmt.Println("Choose item: ")
		fmt.Scan(&choice)
	}
	if equipmentItems[c.contents[choice]] != nil {
		loot := equipmentItems[c.contents[choice]]
		player.inventory = append(player.inventory, loot)
		fmt.Printf("You receive a %s! You put the %s in your backpack.", loot.name, loot.name)
	}
	if defences[c.contents[choice]] != nil {
		player.shield = defences[c.contents[choice]]
		fmt.Printf("Shield upgraded to %s!\n", player.shield.name)
	}
	if weaponry[c.contents[choice]] != nil {
		player.weapon1 = weaponry[c.contents[choice]]
		fmt.Printf("Weapon upgraded to %s!\n", player.weapon1.name)
	}
}

func (c *chest) takeItem(item int) {
	s := c.contents
	s = append(s[:item], s[item+1:]...)
	c.contents = s
}
