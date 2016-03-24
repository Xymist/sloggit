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

func main() {
	//systray.Run(onReady)
	g := &game{Welcome: introductionText, playerCharacter: generateCharacter()}
	g.Play()
}
