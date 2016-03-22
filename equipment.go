package main

// 'Equipment' is all things that can be carried by an organism, which are NOT a weapon or shield

var equipmentItems = map[string]*equipment{
	"Empty":         {physObject: physObject{name: "Empty", integrity: 0}, abilities: []string{}},
	"Torch":         {physObject: physObject{name: "Torch", integrity: 100}, abilities: []string{"Weak Light"}},
	"Climbing Rope": {physObject: physObject{name: "Climbing Rope", integrity: 100}, abilities: []string{"Descend"}},
	"Tinderbox":     {physObject: physObject{name: "Tinderbox", integrity: 100}, abilities: []string{"Ignite"}},
}
