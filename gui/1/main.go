package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 255}
)

func makeSign() fyne.CanvasObject {

	bg := canvas.NewCircle(color.RGBA{255, 0, 0, 255})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5

	bar := canvas.NewRectangle(white)
	c := fyne.NewContainerWithoutLayout(bg, bar)
	bg.Resize(fyne.NewSize(100, 100))
	bg.Move(fyne.NewPos(10, 10))

	bar.Resize(fyne.NewSize(80, 20))
	bar.Move(fyne.NewPos(20, 50))

	return c

}

func main() {
	root := app.New()
	window := root.NewWindow("My program")

	//window.SetContent(canvas.NewText("BICHITA MIA", color.RGBA{230, 0, 0, 1}))
	window.SetContent(makeSign())

	window.SetPadded(false)
	window.Resize(fyne.NewSize(120, 120))

	window.ShowAndRun()
}
