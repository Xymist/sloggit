package main

// 'Equipment' is all things that can be carried by an organism, which are NOT a weapon or shield

type equipment struct {
	physObject
	abilities []string
}

var equipmentItems = map[string]*equipment{
	"Empty":         {physObject: physObject{name: "Empty", integrity: 0}, abilities: []string{}},
	"Torch":         {physObject: physObject{name: "Torch", integrity: 100}, abilities: []string{"Weak Light"}},
	"Climbing Rope": {physObject: physObject{name: "Climbing Rope", integrity: 100}, abilities: []string{"Descend"}},
	"Tinderbox":     {physObject: physObject{name: "Tinderbox", integrity: 100}, abilities: []string{"Ignite"}},
	"Acorn":         {physObject: physObject{name: "Acorn", integrity: 100}, abilities: []string{}},
	"Gold Coin":     {physObject: physObject{name: "Gold Coin", integrity: 100}, abilities: []string{}},
	"Shiny Key":     {physObject: physObject{name: "Shiny Key", integrity: 100}, abilities: []string{}},
}
