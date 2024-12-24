package uimodel

import "fyne.io/fyne/v2"

type Navigations []Navigation

func (navigations Navigations) IDs() []string {
	ids := make([]string, 0, len(navigations))
	for _, navigation := range navigations {
		ids = append(ids, navigation.ID)
	}
	return ids
}

type Navigation struct {
	ID    string
	Title string
	View  func(w fyne.Window) fyne.CanvasObject
}
