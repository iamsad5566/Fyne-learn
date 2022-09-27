package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		// create new window
		w := app.NewWindow(
			app.Title("Egg timer"),
			app.Size(unit.Dp(1400), unit.Dp(800)),
		)

		// ops are the operations from the UI
		var ops op.Ops

		// startButton is a clickable widget
		var startButton widget.Clickable

		// th defnes the material design style
		th := material.NewTheme(gofont.Collection())

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
	for e := range w.Events() {

		// detect what type of event
		switch e := e.(type) {

		// this is sent when the application should re-render.
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			// Let's try out the flexbox layout concept:
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// We insert two rigid elements:
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
								btn := material.Button(th, &startButton, "Start")
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

	}
	return nil
}
