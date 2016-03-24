package main

import (
	"math/rand"
	"time"

	"github.com/getlantern/systray"
)

func genRandom(min, max int) int {
	rSource := rand.NewSource(time.Now().UnixNano())
	rRand := rand.New(rSource)
	return rRand.Intn(max-min) + min
}

func main() {
	systray.Run(onReady)
}
