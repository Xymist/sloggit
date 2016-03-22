package main

var defences = map[string]*shield{
	"Battered Shield": {physObject: physObject{name: "Battered Shield", integrity: 6}, bonuses: []string{"No Bonus"}},
	"Wooden Shield":   {physObject: physObject{name: "Wooden Shield", integrity: 12}, bonuses: []string{"No Bonus"}},
}
