package main

import (
    _ "fyne.io/fyne/v2"  // add this line to your code
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("Hello World")

    w.SetContent(widget.NewLabel("Hello World!"))
    w.ShowAndRun()
}


