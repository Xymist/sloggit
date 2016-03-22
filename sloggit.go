package main

func main() {
	var introduction = "You wake up in a clearing in the forest, resting on a thick carpet of pine needles. \nYou stand up and look about you, your hand instinctively reaching \nfor your Rusty Sword and Battered Shield. There is a chill in the air; \nfrom that and the trees you deduce that you must be many hundreds of miles \nfurther North than where you laid down to sleep. Suspecting foul play, \nyou check that you are undamaged and set off through the trees...\n\n"

	g := &Game{Welcome: introduction, playerCharacter: generateCharacter()}
	g.Play()
}
