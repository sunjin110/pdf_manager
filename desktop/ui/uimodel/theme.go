package uimodel

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type Theme struct {
	fyne.Theme
	Variant fyne.ThemeVariant
}

func (f *Theme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.Variant)
}
