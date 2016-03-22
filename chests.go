package main

// 'Chests' contains all objects which contain equipment, whether it's a jar, a bookcase, a desk or actually a chest.

var chests = map[string]*chest{
	"Cave Chest": {physObject: physObject{name: "Cave Chest", integrity: 100}, contents: []string{"Wooden Shield", "Torch", "Climbing Rope", "Tinderbox"}},
}
