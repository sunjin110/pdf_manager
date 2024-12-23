package main

import (
	"fmt"
	"log"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sunjin110/pdf_manager/core"
	"github.com/sunjin110/pdf_manager/desktop/ui/uidata"
	"github.com/sunjin110/pdf_manager/desktop/ui/uimodel"
)

func main() {
	a := app.NewWithID("info.sunjin.pdf_manager_desktop")

	w := a.NewWindow("🐼 PDF Manager")

	pdfManagerCore, err := core.NewCore(filepath.Join(a.Storage().RootURI().Path(), "app.db"))
	if err != nil {
		log.Fatalf("failed new core. err: %v", err)
	}
	fmt.Println("sqlite path is ", filepath.Join(a.Storage().RootURI().Path(), "app.db"))

	a.Settings().SetTheme(uidata.DarkTheme)
	title := widget.NewLabel("タイトル")
	content := container.NewStack()
	splitContent := container.NewHSplit(makeNav(w, pdfManagerCore, uidata.Navigations, title, content), content)
	splitContent.SetOffset(0.3)
	w.SetContent(splitContent)

	w.Resize(fyne.NewSize(760, 580))
	w.ShowAndRun()
}

func makeNav(parentWindow fyne.Window, pdfManagerCore core.Core, navigations uimodel.Navigations, title *widget.Label, content *fyne.Container) fyne.CanvasObject {
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
		content.Objects = []fyne.CanvasObject{navigation.View(parentWindow, pdfManagerCore)}
	}
	return tree
}
