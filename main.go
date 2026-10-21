package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("vespa-term")
	w.SetContent(widget.NewLabel("vespa-term"))
	w.ShowAndRun()
}
