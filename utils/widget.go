package utils

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var Cfg Config

func (c *Config) MakeUi(w fyne.Window) (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	c.EditWidget = edit
	c.PreviewWidget = preview

	edit.OnChanged = func(content string) {
		preview.ParseMarkdown(content)
		c.SaveMenuItem.Disabled = false
	}

	openMenuItem := fyne.NewMenuItem("Open", func() {
		dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if r == nil {
				return
			}
			data, err := ioutil.ReadAll(r)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			c.CurrentFile = r.URI()
			edit.SetText(string(data))
			c.SaveMenuItem.Disabled = true
			w.SetTitle("Markdown Notepad - " + c.CurrentFile.Name())
		}, w).Show()
	})

	saveMenuItem := fyne.NewMenuItem("Save", func() {
		if c.CurrentFile == nil {
			dialog.NewFileSave(func(uri fyne.URIWriteCloser, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if uri == nil {
					return
				}
				c.CurrentFile = uri.URI()
				writeData := []byte(edit.Text)
				if _, err := uri.Write(writeData); err != nil {
					dialog.ShowError(err, w)
					return
				}
				uri.Close()
				c.SaveMenuItem.Disabled = true
				w.SetTitle("Markdown Notepad - " + c.CurrentFile.Name())
			}, w).Show()
		} else {
			writeData := []byte(edit.Text)
			if err := ioutil.WriteFile(c.CurrentFile.Path(), writeData, 0644); err != nil {
				dialog.ShowError(err, w)
				return
			}
			c.SaveMenuItem.Disabled = true
			w.SetTitle("Markdown Notepad - " + c.CurrentFile.Name())
		}
	})

	c.SaveMenuItem = saveMenuItem
	saveMenuItem.Disabled = true

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem)
	mainMenu := fyne.NewMainMenu(fileMenu)

	w.SetMainMenu(mainMenu)

	return edit, preview
}
