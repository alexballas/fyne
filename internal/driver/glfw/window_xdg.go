//go:build linux || freebsd || openbsd || netbsd

package glfw

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/internal/build"
	"fyne.io/fyne/v2/storage"

	"github.com/godbus/dbus/v5"
)

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}

func (w *window) handleDrop(names []string) []fyne.URI {
	var (
		uris []fyne.URI
		conn *dbus.Conn
		err  error
	)

	if build.IsFlatpak {
		conn, err = dbus.SessionBus()
	}

	for _, name := range names {
		matched := false
		if conn != nil && err == nil {
			obj := conn.Object("org.freedesktop.portal.Desktop", "/org/freedesktop/portal/desktop")
			call := obj.Call("org.freedesktop.portal.FileTransfer.RetrieveFiles", 0, name, map[string]dbus.Variant{})
			if call.Err == nil {
				var files []string
				if err := call.Store(&files); err == nil {
					for _, f := range files {
						uris = append(uris, storage.NewFileURI(f))
					}
					matched = true
				}
			}
		}

		if !matched {
			uris = append(uris, storage.NewFileURI(name))
		}
	}
	return uris
}
