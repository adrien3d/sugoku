package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

/*
	Libraries to use

	https://github.com/getlantern/systray
	https://github.com/shurcooL/trayhost
*/

func main() {
	grid := make([][]int, 9)

	for i := 0; i < 9; i++ {
		grid[i] = make([]int, 9)
	}
	fmt.Println("grid: ", grid)

	err := ui.Main(func() {
		button := ui.NewButton("Validate")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("The grid:"), true)
		box.Append(ui.NewEntry(), true)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Sugoku", 600, 600, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, !")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}

}
