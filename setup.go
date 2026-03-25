package main

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed icon.ico
var resourceIconIcoData []byte

var resourceIconIco *fyne.StaticResource

func init() {
	resourceIconIco = &fyne.StaticResource{
		StaticName:    "icon.ico",
		StaticContent: resourceIconIcoData,
	}
}
