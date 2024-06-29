package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/NikoMalik/Markdown-notepad/utils"
)

func main() {
	// create fyne app
	a := app.New()

	// create window for app
	win := a.NewWindow("Markdown Notepad")
	win.Resize(fyne.NewSize(800, 800))

	// get the user interface for notepad
	edit, preview := utils.Cfg.MakeUi(win)

	// create a menu for changing the theme
	themeMenu := fyne.NewMenu("Themes",
		fyne.NewMenuItem("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
		fyne.NewMenuItem("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
	)

	// add the theme menu to the main menu
	mainMenu := fyne.NewMainMenu(themeMenu)
	win.SetMainMenu(mainMenu)

	// set the content of the window
	content := container.NewHSplit(edit, preview)
	content.SetOffset(0.5)
	win.SetContent(content)

	// show window
	win.ShowAndRun()
}
