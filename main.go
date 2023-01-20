package main

import (
	"books-app/books"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type ui struct {
	currentPage *books.Page
	pages       books.Pages
	content     *fyne.Container
	tools       toolbar
}

type toolbar struct {
	plus  *widget.Toolbar
	minus *widget.Toolbar
}

func (u *ui) nextPage() {
	for i, v := range u.pages {
		var nextPage books.Page
		thisPage := v

		if i == len(u.pages)-1 {
			break
		}

		if thisPage.Text == u.currentPage.Text {
			nextPage = u.pages[i+1]
			u.currentPage = &nextPage
			break
		}
	}
	u.content = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, u.tools.minus, u.tools.plus), u.tools.minus, u.newContent(), u.tools.plus)
}

func (u *ui) lastPage() {
	for i, v := range u.pages {
		var lastPage books.Page
		thisPage := v

		if thisPage.Text == u.currentPage.Text {
			if (i - 1) < 0 {
				break
			}
			lastPage = u.pages[i-1]
			u.currentPage = &lastPage
			break
		}
	}
	u.content = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, u.tools.minus, u.tools.plus), u.tools.minus, u.newContent(), u.tools.plus)
}

func (u *ui) newContent() *container.Split {
	top := canvas.NewImageFromFile(u.currentPage.Image)
	bottom := widget.NewLabel(u.currentPage.Text)

	bottom.Alignment = fyne.TextAlignCenter
	bottom.Wrapping = fyne.TextWrapBreak

	content := container.NewVSplit(
		top,
		bottom,
	)

	content.Offset = 0.75

	return content
}

func main() {
	a := app.New()
	w := a.NewWindow("book")

	ui := ui{
		currentPage: &books.Books[0].Pages[0],
		pages:       books.Books[0].Pages,
	}

	top := canvas.NewImageFromFile(ui.currentPage.Image)
	bottom := widget.NewLabel(ui.currentPage.Text)

	bottom.Alignment = fyne.TextAlignCenter
	bottom.Wrapping = fyne.TextWrapBreak

	content := container.NewVSplit(
		top,
		bottom,
	)

	content.Offset = 0.75

	plusBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			ui.nextPage()
			w.SetContent(ui.content)
		}),
	)

	minusBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			ui.lastPage()
			w.SetContent(ui.content)
		}),
	)

	ui.tools = toolbar{
		plusBar,
		minusBar,
	}

	ui.content = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, minusBar, plusBar), minusBar, content, plusBar)

	w.SetContent(ui.content)

	w.Resize(fyne.NewSize(600, 800))

	w.Show()
	a.Run()
}
