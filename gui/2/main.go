package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	root := app.NewWindow("The title")
	root.ShowAndRun()
}
