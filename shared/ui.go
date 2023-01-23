package shared

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func (u *UI) NextPage() {
	for i, v := range u.Pages {
		var nextPage Page
		thisPage := v

		if i == len(u.Pages)-1 {
			break
		}

		if thisPage.Text == u.CurrentPage.Text {
			nextPage = u.Pages[i+1]
			u.CurrentPage = &nextPage
			break
		}
	}
	u.Content = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, u.Tools.Minus, u.Tools.Plus), u.Tools.Minus, u.newContent(), u.Tools.Plus)
}

func (u *UI) LastPage() {
	for i, v := range u.Pages {
		var lastPage Page
		thisPage := v

		if thisPage.Text == u.CurrentPage.Text {
			if (i - 1) < 0 {
				break
			}
			lastPage = u.Pages[i-1]
			u.CurrentPage = &lastPage
			break
		}
	}
	u.Content = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, u.Tools.Minus, u.Tools.Plus), u.Tools.Minus, u.newContent(), u.Tools.Plus)
}

func (u *UI) newContent() *container.Split {
	top := canvas.NewImageFromFile(u.CurrentPage.Image)
	bottom := widget.NewLabel(u.CurrentPage.Text)

	bottom.Alignment = fyne.TextAlignCenter
	bottom.Wrapping = fyne.TextWrapBreak

	content := container.NewVSplit(
		top,
		bottom,
	)

	content.Offset = 0.75

	return content
}

func (u *UI) LoadUI(w fyne.Window, b map[string]Book, a fyne.App) {
	w = WelcomeWindow(a, b)
	w.Show()
}

func WelcomeWindow(app fyne.App, b map[string]Book) fyne.Window {
	w := app.NewWindow("main menu")

	contain := container.NewGridWithRows(2,
		widget.NewButton("New Book", func() {
			w.Hide()
		}),
		widget.NewButton("Current Books", func() {
			CurrentBooksMenu(w, b)
		}),
	)
	w.SetContent(contain)
	w.Resize(fyne.NewSize(600, 800))

	return w
}

func CurrentBooksMenu(w fyne.Window, bookList map[string]Book) {

	list := widget.Box{Children: nil}

	for k, v := range bookList {
		list.Append(widget.NewButton(k, func() {
			BookViewer(w, bookList, v.Title)
		}))
	}

	side := fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1), &list)

	w.SetContent(side)
}

func BookViewer(w fyne.Window, b map[string]Book, name string) {

	u := UI{
		CurrentPage: &b[name].Pages[0],
		Pages:       b[name].Pages,
	}
	top := canvas.NewImageFromFile(u.CurrentPage.Image)
	bottom := widget.NewLabel(u.CurrentPage.Text)

	bottom.Alignment = fyne.TextAlignCenter
	bottom.Wrapping = fyne.TextWrapBreak

	content := container.NewVSplit(
		top,
		bottom,
	)

	content.Offset = 0.75

	plusBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			u.NextPage()
			w.SetContent(u.Content)
		}),
	)

	minusBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			u.LastPage()
			w.SetContent(u.Content)
		}),
	)

	u.Tools = Toolbar{
		plusBar,
		minusBar,
	}

	u.Content = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, minusBar, plusBar), minusBar, content, plusBar)

	w.SetContent(u.Content)
}
