// package main

// import (
// 	"fmt"
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// 	"github.com/atotto/clipboard"
// )

// var emojis = []string{"😀", "😂", "❤️", "🔥", "🚀", "🎉", "👍", "🐱"}

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Emoji Picker")
// 	w.Resize(fyne.NewSize(300, 400))

// 	list := widget.NewList(
// 		func() int { return len(emojis) },
// 		func() fyne.CanvasObject { return widget.NewLabel("") },
// 		func(i widget.ListItemID, obj fyne.CanvasObject) {
// 			obj.(*widget.Label).SetText(emojis[i])
// 		},
// 	)

// 	list.OnSelected = func(id widget.ListItemID) {
// 		emoji := emojis[id]
// 		clipboard.WriteAll(emoji)
// 		fmt.Println("Copiado:", emoji)
// 	}

// 	w.SetContent(container.NewVBox(list))
// 	w.ShowAndRun()
// }

package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mw *walk.MainWindow
	MainWindow{
		AssignTo: &mw,
		Title:    "Walk Test",
		Size:     Size{Width: 300, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			Label{Text: "Olá, Walk!"},
			PushButton{
				Text: "Fechar",
				OnClicked: func() {
					mw.Close()
				},
			},
		},
	}.Run()
}
