package main

type physObject struct {
	name      string
	integrity int
}

func (p *physObject) takeDamage(damage int) {
	p.integrity -= damage
}
