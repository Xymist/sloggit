package main

//This works perfectly well on Linux, but for whatever reason totally fails to compile for OSX and breaks the build for Windows.
//On hold until I have an answer from the package maintainer.

/*import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/skratchdot/open-golang/open"
)

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Sloggit")
	systray.SetTooltip("A little text adventure")
	mQuit := systray.AddMenuItem("Quit", "Exit Sloggit")
	mURL := systray.AddMenuItem("Source", "Go To Github")
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quitting now...")
				os.Exit(0)
			case <-mURL.ClickedCh:
				open.Run("https://github.com/Xymist/sloggit")
			}
		}
	}()

	go func() {
		g := &game{Welcome: introductionText, playerCharacter: generateCharacter()}
		g.Play()
		os.Exit(0)
	}()
}
*/
