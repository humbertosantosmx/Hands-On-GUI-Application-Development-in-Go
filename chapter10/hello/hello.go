package main

import "github.com/fyne-io/fyne/widget"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	w := app.NewWindow("Hello World")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
