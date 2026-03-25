package main

import (
	"fmt"
	"image/color"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("WinShaper")
	w := a.NewWindow("WinShaper")
	w.SetIcon(resourceIconIco)
	w.Resize(fyne.NewSize(525, 550))
	w.SetFixedSize(true)

	makeTitle := func(text string) *canvas.Text {
		t := canvas.NewText(text, color.White)
		t.TextStyle = fyne.TextStyle{Bold: true}
		t.TextSize = 15
		return t
	}

	// --- [ UI ELEMENTS ] ---

	label := widget.NewLabel("WinShaper")
	label.TextStyle.Bold = true
	label.Alignment = fyne.TextAlignCenter

	checkDarkMode := widget.NewCheck("Enable Dark Mode", func(value bool) {
		val := "1"
		if value {
			val = "0"
		}
		exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, "/v", "AppsUseLightTheme", "/t", "REG_DWORD", "/d", val, "/f").Run()
		exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`, "/v", "SystemUsesLightTheme", "/t", "REG_DWORD", "/d", val, "/f").Run()
	})
	checkDarkMode.Checked = isDarkModeEnabled()

	checkOldMenu := widget.NewCheck("Classic Context Menu", func(value bool) {
		if value {
			exec.Command("reg", "add", `HKCU\SOFTWARE\CLASSES\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}\InprocServer32`, "/ve", "/f").Run()
			exec.Command("reg", "add", `HKCU\Software\Classes\CLSID\{d93ed569-3b3e-4bff-8355-3c44f6a52bb5}\InprocServer32`, "/ve", "/f").Run()
		} else {
			exec.Command("reg", "delete", `HKCU\Software\Classes\CLSID\{86ca1aa0-34aa-4e8b-a509-50c905bae2a2}`, "/f").Run()
			exec.Command("reg", "delete", `HKCU\Software\Classes\CLSID\{d93ed569-3b3e-4bff-8355-3c44f6a52bb5}`, "/f").Run()
		}
	})
	checkOldMenu.Checked = isOldMenuEnabled()

	selectionColorPicker := dialog.NewColorPicker("Choose Selection Color", "", func(c color.Color) {
		r, g, b, _ := c.RGBA()
		rgbString := fmt.Sprintf("%d %d %d", r>>8, g>>8, b>>8)
		exec.Command("reg", "add", `HKCU\Control Panel\Colors`, "/v", "Hilight", "/t", "REG_SZ", "/d", rgbString, "/f").Run()
		exec.Command("reg", "add", `HKCU\Control Panel\Colors`, "/v", "HotTrackingColor", "/t", "REG_SZ", "/d", rgbString, "/f").Run()
	}, w)
	selectionColorPicker.Advanced = true

	btnSelectionColor := widget.NewButton("Change Selection Color (Restart Required)", func() {
		selectionColorPicker.Show()
	})

	checkArrow := widget.NewCheck("Remove Shortcut Arrows", func(value bool) {
		if value {
			exec.Command("reg", "add", `HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Shell Icons`, "/v", "29", "/t", "REG_SZ", "/d", `%windir%\System32\shell32.dll,51`, "/f").Run()
		} else {
			exec.Command("reg", "delete", `HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Shell Icons`, "/v", "29", "/f").Run()
		}
	})
	checkArrow.Checked = isArrowEnabled()

	checkBin := widget.NewCheck("Hide Recycle Bin on Desktop", func(value bool) {
		val := "0"
		if value {
			val = "1"
		}
		exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\HideDesktopIcons\NewStartPanel`, "/v", "{645FF040-5081-101B-9F08-00AA002F954E}", "/t", "REG_DWORD", "/d", val, "/f").Run()
	})
	checkBin.Checked = isRecycleBinHidden()

	checkHidden := widget.NewCheck("Show Hidden Files", func(value bool) {
		val := "2"
		if value {
			val = "1"
		}
		exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, "/v", "Hidden", "/t", "REG_DWORD", "/d", val, "/f").Run()
	})
	checkHidden.Checked = isHiddenFilesShown()

	checkExt := widget.NewCheck("Show File Extensions", func(value bool) {
		val := "1"
		if value {
			val = "0"
		}
		exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, "/v", "HideFileExt", "/t", "REG_DWORD", "/d", val, "/f").Run()
	})
	checkExt.Checked = isExtensionsShown()

	checkSeconds := widget.NewCheck("Show Seconds in System Clock", func(value bool) {
		val := "0"
		if value {
			val = "1"
		}
		exec.Command("reg", "add", `HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, "/v", "ShowSecondsInSystemClock", "/t", "REG_DWORD", "/d", val, "/f").Run()
	})
	checkSeconds.Checked = isSecondsEnabled()

	buttonApply := widget.NewButton("Apply Changes", func() {
		exec.Command("taskkill", "/f", "/im", "explorer.exe").Run()
		exec.Command("cmd", "/c", "start explorer.exe").Run()
	})
	buttonApply.Importance = widget.HighImportance

	// --- [ LAYOUT ] ---

	canvasLayout := container.NewVBox(
		widget.NewSeparator(),
		makeTitle("System UI"),
		checkDarkMode,
		checkOldMenu,
		btnSelectionColor,
		makeTitle("Desktop"),
		checkArrow,
		checkBin,
		makeTitle("Explorer"),
		checkHidden,
		checkExt,
		makeTitle("Taskbar"),
		checkSeconds,
	)

	w.SetContent(container.NewBorder(
		label,
		buttonApply,
		nil,
		nil,
		container.NewPadded(canvasLayout),
	))

	w.ShowAndRun()
}
