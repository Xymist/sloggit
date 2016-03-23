package main

type hittable interface {
	takeDamage(int)
	updateStatus(*statusEffect)
}

type entity interface {
	inspect()
}

func examine(e entity) {
	e.inspect()
}
