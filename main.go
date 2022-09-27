package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var progress float32
var progressIncrementer chan float32
var boiling bool

type Point struct {
	X, Y float32
}

type Rectangle struct {
	Min, Max Point
}

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Egg timer"),
			app.Size(unit.Dp(400), unit.Dp(600)),
		)

		// ops are the operations from the UI
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defnes the material design style
		th := material.NewTheme(gofont.Collection())

		progressIncrementer = make(chan float32)
		go func() {
			for {
				time.Sleep(time.Second / 25)
				progressIncrementer <- 0.004
			}
		}()

		if err := draw(w, ops, startButton, th); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window, ops op.Ops, startButton widget.Clickable, th *material.Theme) error {
	type C = layout.Context
	type D = layout.Dimensions
	// listen for events in the window.
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			// this is sent when the application should re-render.
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				if startButton.Clicked() {
					boiling = !boiling
				}

				// Let's try out the flexbox layout concept:
				layout.Flex{
					// Vertical alignment, from top to bottom
					Axis: layout.Vertical,
					// Empty space is left at the start, i.e. at the top
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					// We insert two rigid elements:
					layout.Rigid(
						func(gtx C) D {
							circle := clip.Ellipse{
								// Hard coding the x coordinate. Try resizing the window
								// Min: image.Pt(80, 0),
								// Max: image.Pt(320, 240),
								// Soft coding the x coordinate. Try resizing the window
								Min: image.Pt(gtx.Constraints.Max.X/2-300, 0),
								Max: image.Pt(gtx.Constraints.Max.X/2+300, 600),
							}.Op(gtx.Ops)
							color := color.NRGBA{R: 200, A: 255}
							paint.FillShape(gtx.Ops, color, circle)
							d := image.Point{Y: 800}
							return layout.Dimensions{Size: d}
						},
					),

					layout.Rigid(
						func(gtx C) D {
							bar := material.ProgressBar(th, progress)
							return bar.Layout(gtx)
						},
					),
					// First a button ...
					layout.Rigid(
						func(gtx C) D {
							margins := layout.Inset{
								Top:    unit.Dp(25),
								Bottom: unit.Dp(25),
								Left:   unit.Dp(25),
								Right:  unit.Dp(25),
							}

							return margins.Layout(gtx,
								func(gtx layout.Context) D {
									text := ""
									if boiling {
										text = "Stop"
									} else {
										text = "Start"
									}
									btn := material.Button(th, &startButton, text)
									return btn.Layout(gtx)
								},
							)
						},
					),

					// ... then an empty spacer
					layout.Rigid(
						// The height of the spacer is 25 Device independent pixels
						layout.Spacer{Height: unit.Dp(25)}.Layout,
					),
				)
				e.Frame(gtx.Ops)

			case system.DestroyEvent:
				return e.Err
			}

		case p := <-progressIncrementer:
			if boiling && progress < 1 {
				progress += p
				w.Invalidate()
			}
		}
	}
}
