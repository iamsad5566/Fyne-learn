package clock

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Clock() {
	app := app.New()
	window := app.NewWindow("Tik Tok")
	clock := widget.NewLabel("")
	updateTime(clock)
	window.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	window.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
