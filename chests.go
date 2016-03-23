package main

import "fmt"

// 'Chests' contains all objects which contain equipment, whether it's a jar, a bookcase, a desk or actually a chest.

var chests = map[string]*chest{
	"Cave Chest": {physObject: physObject{name: "Cave Chest", integrity: 100}, contents: []string{"Wooden Shield", "Torch", "Climbing Rope", "Tinderbox"}},
}

func open(c chest) *equipment {
	fmt.Println("You see something inside:")
	for index, element := range c.contents {
		fmt.Printf("\t%d - %s\n", index+1, element)
	}
	choice := 0
	for choice < 1 || choice > len(c.contents) { //Loop until they pick something within bounds
		fmt.Println("Choose item: ")
		fmt.Scan(&choice)
	}
	loot := equipmentItems[c.contents[choice]]
	return loot
}
