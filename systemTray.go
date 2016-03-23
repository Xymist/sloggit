package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Sloggit")
	systray.SetTooltip("A little text adventure")
	mQuit := systray.AddMenuItem("Quit", "Exit Sloggit")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		fmt.Println("Quitting now...")
		os.Exit(0)
	}()

	go func() {
		var introduction = "You wake up in a clearing in the forest, resting on a thick carpet of pine needles. You stand up and look about you, your hand instinctively reaching for your Rusty Sword and Battered Shield. There is a chill in the air; from that and the trees you deduce that you must be many hundreds of miles further North than where you laid down to sleep. Suspecting foul play, you check that you are undamaged and set off through the trees...\n\n"

		g := &Game{Welcome: introduction, playerCharacter: generateCharacter()}
		g.Play()
	}()
}
