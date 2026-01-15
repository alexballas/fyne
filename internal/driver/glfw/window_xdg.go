//go:build linux || freebsd || openbsd || netbsd

package glfw

import (
	"github.com/godbus/dbus/v5"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/build"
	"fyne.io/fyne/v2/storage"
)

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}

func (w *window) handleDrop(names []string) []fyne.URI {
	if build.IsFlatpak && len(names) == 1 {
		conn, err := dbus.SessionBus()
		if err == nil {
			obj := conn.Object("org.freedesktop.portal.Documents", "/org/freedesktop/portal/documents")
			call := obj.Call("org.freedesktop.portal.FileTransfer.RetrieveFiles", 0, names[0], map[string]dbus.Variant{})
			if call.Err == nil {
				var files []string
				if err := call.Store(&files); err == nil {
					uris := make([]fyne.URI, len(files))
					for i, f := range files {
						uris[i] = storage.NewFileURI(f)
					}
					return uris
				}
			}
		}
	}

	uris := make([]fyne.URI, len(names))
	for i, name := range names {
		uris[i] = storage.NewFileURI(name)
	}
	return uris
}
