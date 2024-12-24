package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sunjin110/pdf_manager/desktop/ui/uidata"
	"github.com/sunjin110/pdf_manager/desktop/ui/uimodel"
)

func main() {
	fmt.Println("hello world")

	a := app.New()
	w := a.NewWindow("🐼 PDF Manager")

	a.Settings().SetTheme(uidata.DarkTheme)

	title := widget.NewLabel("タイトル")

	content := container.NewStack()

	mainContainer := container.NewBorder(
		title, nil, nil, content,
	)

	splitContent := container.NewHSplit(makeNav(uidata.Navigations, title), mainContainer)
	splitContent.SetOffset(0.3)
	w.SetContent(splitContent)

	w.Resize(fyne.NewSize(760, 580))
	w.ShowAndRun()
}

func makeNav(navigations uimodel.Navigations, title *widget.Label) fyne.CanvasObject {
	navigationMap := make(map[string]uimodel.Navigation, len(navigations))
	for _, navigation := range navigations {
		navigationMap[navigation.ID] = navigation
	}

	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return navigations.IDs()
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == ""
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			navigation := navigationMap[id]
			o.(*widget.Label).SetText(navigation.Title)
		},
	)

	tree.OnSelected = func(id widget.TreeNodeID) {
		navigation := navigationMap[id]
		title.SetText(navigation.Title)
	}
	return tree
}
