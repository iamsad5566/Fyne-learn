package main

import (
	"image/color"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	rect := canvas.NewRectangle(blue)
	myCanvas.SetContent(rect)

	go func() {
		time.Sleep(time.Second)
		green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
		rect.FillColor = green
		setContentToText(myCanvas)
		rect.Refresh()
		time.Sleep(time.Second * 2)
		rect.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		setContentToCircle(myCanvas)

		// widgt := widget.NewLabel("test")
		a := canvas.NewCircle(color.Black)
		b := canvas.NewText("Test", green)
		input := widget.NewEntry()
		input.SetPlaceHolder("Enter text...")
		c := container.NewVBox(input, widget.NewButton("Save", func() {
			log.Println("Content was:", input.Text)
		}))
		content := container.NewGridWithColumns(2, a, b, c)
		myCanvas.SetContent(content)

		rect.Refresh()
	}()

	myWindow.Resize(fyne.NewSize(100, 100))
	myWindow.ShowAndRun()
}

func setContentToText(c fyne.Canvas) {
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Text", green)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red
	c.SetContent(circle)
}
