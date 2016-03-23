package main

import (
	"math/rand"
	"time"
)

func genRandom(min, max int) int {
	rSource := rand.NewSource(time.Now().UnixNano())
	rRand := rand.New(rSource)
	return rRand.Intn(max-min) + min
}

type hittable interface {
	takeDamage(int)
}

type entity interface {
	inspect()
}

func attack(h hittable, weapon weapon) {
	h.takeDamage(weapon.baseDamage)
	weapon.integrity--
}

func examine(e entity) {
	e.inspect()
}
