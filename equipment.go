package main

// 'Equipment' is all things that can be carried by an organism, which are NOT a weapon or shield

var equipmentItems = map[string]*equipment{
	"Empty":         {physObject: physObject{name: "Empty", integrity: 0}},
	"Torch":         {physObject: physObject{name: "Torch", integrity: 100}},
	"Climbing Rope": {physObject: physObject{name: "Climbing Rope", integrity: 100}},
	"Tinderbox":     {physObject: physObject{name: "Tinderbox", integrity: 100}},
}
