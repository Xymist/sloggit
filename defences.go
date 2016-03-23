package main

import "fmt"

//Shields are anything you hide behind and carry, regardless of material or magicness

type shield struct {
	physObject
	bonuses []string
}

func (s *shield) inspect() {
	fmt.Println(s.name, "Integrity: ", s.integrity)
}

var defences = map[string]*shield{
	"Battered Shield": {physObject: physObject{name: "Battered Shield", integrity: 6}, bonuses: []string{"No Bonus"}},
	"Wooden Shield":   {physObject: physObject{name: "Wooden Shield", integrity: 12}, bonuses: []string{"No Bonus"}},
}
