package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	root := app.NewWindow("The title")
	root.SetContent(widget.NewButton("HIT ME", func() { println("ASDASDASD") }))
	root.ShowAndRun()
}
