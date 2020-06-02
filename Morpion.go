package main

import (
	"strconv"

	"github.com/tadvi/winc"
)

func main() {

	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(600, 600)
	mainWindow.SetText("Jeu Du Morpion")

	var Tour = "X"
	var Text = "C'est au tour de : " + Tour
	var PosX = 160
	var PosY = 160

	txt := winc.NewLabel(mainWindow)
	txt.SetText(Text)
	txt.SetSize(1000, 20)

	var numCase = 0

	tab := make(map[string]string)

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			numCase = numCase + 1
			newButton := winc.NewPushButton(mainWindow)
			newButton.SetSize(80, 80)
			newButton.SetPos(PosY+y*80, PosX+x*80)
			name := strconv.Itoa(numCase)
			newButton.SetText(name)

			newButton.OnClick().Bind(func(e *winc.Event) {
				if tab[name] == "" {
					newButton.SetText(Tour)
					tab[name] = Tour

					if Tour == "X" {
						Tour = "O"
					} else {
						Tour = "X"
					}

					Text = "Tour: " + Tour
					txt.SetText(Text)
				}
			})
		}
	}

	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
