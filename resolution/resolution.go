package resolution

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var width float32 = 400
var height float32 = 600

func Resolution() {
	app := app.New()
	window := app.NewWindow("resolution")
	window.Resize(fyne.NewSize(width, height))
	wInput := widget.NewEntry()
	wInput.SetPlaceHolder("width")
	hInput := widget.NewEntry()
	hInput.SetPlaceHolder("height")

	content := container.NewVBox(wInput, hInput, widget.NewButton("Save", func() {
		w, _ := strconv.ParseFloat(wInput.Text, 32)
		h, _ := strconv.ParseFloat(hInput.Text, 32)
		width = float32(w)
		height = float32(h)
		window.Resize(fyne.NewSize(width, height))
	}))
	window.SetContent(content)
	window.ShowAndRun()
}
