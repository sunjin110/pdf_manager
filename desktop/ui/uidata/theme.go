package uidata

import (
	"fyne.io/fyne/v2/theme"
	"github.com/sunjin110/pdf_manager/desktop/ui/uimodel"
)

var (
	// DarkTheme ダークモード
	DarkTheme = &uimodel.Theme{
		Variant: theme.VariantDark,
		Theme:   theme.DefaultTheme(),
	}

	LightTheme = &uimodel.Theme{
		Variant: theme.VariantLight,
		Theme:   theme.DefaultTheme(),
	}
)
