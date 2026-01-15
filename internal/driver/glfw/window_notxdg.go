//go:build !linux && !freebsd && !openbsd && !netbsd

package glfw

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func (w *window) platformResize(canvasSize fyne.Size) {
	d, ok := fyne.CurrentApp().Driver().(*gLDriver)
	if !ok { // don't wait to redraw in this way if we are running on test
		w.canvas.Resize(canvasSize)
		return
	}

	w.canvas.Resize(canvasSize)
	d.repaintWindow(w)
}

func (w *window) handleDrop(names []string) []fyne.URI {
	uris := make([]fyne.URI, len(names))
	for i, name := range names {
		uris[i] = storage.NewFileURI(name)
	}
	return uris
}
