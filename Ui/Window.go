package ui

import (
	"fmt"
	"locksmith/crypter"
	"locksmith/utilities"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EditFile(cipherBytes *[]byte, cipherFile string, key string) {
	a := app.New()
	w := a.NewWindow("Locksmith")
	defer a.Quit()

	textArea := widget.NewMultiLineEntry()
	textArea.Wrapping = fyne.TextWrapWord
	textArea.SetText(string(*cipherBytes))

	saveFile := func() {
		lock, err := crypter.LoadKey(key)
		if err != nil {
			utilities.LogIfError(err)
			return
		}
		new_value := []byte(textArea.Text)
		lock.LoadMessage(new_value)
		if crypter.FindByteDiff(cipherBytes, &new_value) {
			lock.Encrypt(cipherFile)
		} else {
			fmt.Println("No Change Detected!")
		}
		a.Quit()
	}

	menu := fyne.NewMainMenu(
		fyne.NewMenu("Save",
			fyne.NewMenuItem("Save", saveFile),
			fyne.NewMenuItem("Exit", func() { a.Quit() }),
		),
	)

	w.SetMainMenu(menu)
	w.SetContent(container.NewBorder(nil, nil, nil, nil, textArea))
	w.Resize(fyne.NewSize(1024, 1024))
	w.ShowAndRun()
}
