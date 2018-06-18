// Package browser provides Go bindings to browser APIs.  This package is
// intended to be used with WebAssembly to abstract Go developers from having
// to deal with writing javascript inside of Go.
package browser

import (
	"syscall/js"
)

// Window is a browser window.
type Window struct {
	Element
	window         js.Value
	Console        Debug
	Document       Htmldoc
	LocalStorage   WebStorage
	SessionStorage WebStorage
	Path2D         Path2D
}

// GetWindow returns the main browser window object.
func GetWindow() Window {
	jsGlobal := js.Global()
	return Window{
<<<<<<< HEAD
		window:         jsGlobal,
		Console:        Debug{console: jsGlobal.Get("console")},
		Document:       Htmldoc{document: jsGlobal.Get("document")},
		LocalStorage:   WebStorage{storage: jsGlobal.Get("localStorage")},
		SessionStorage: WebStorage{storage: jsGlobal.Get("sessionStorage")},
		Path2D:         Path2D{path2D: jsGlobal.Get("Path2D")},
=======
		window:         js.Global,
		Console:        Debug{console: js.Global.Get("console")},
		Document:       Htmldoc{Element: Element{el: js.Global.Get("document")}},
		LocalStorage:   WebStorage{storage: js.Global.Get("localStorage")},
		SessionStorage: WebStorage{storage: js.Global.Get("sessionStorage")},
		Path2D:         Path2D{path2D: js.Global.Get("Path2D")},
>>>>>>> Update APIs
	}
}

// Alert shows an alert box in the browser with an OK button.
func (w *Window) Alert(m string) {
	w.window.Call("alert", m)
}

// Invoke executes the Callback that is passed in.
func Invoke(cb js.Callback) {
	js.ValueOf(cb).Invoke()
}

// NewWindow creates a new browser window.  This is analogous to Javascript's
// window.open().
func NewWindow(url string) Window {
	jsGlobal := js.Global()
	w := jsGlobal.Call("open", url)

	return Window{
		window:         w,
		Console:        Debug{console: w.Get("console")},
<<<<<<< HEAD
		Document:       Htmldoc{document: w.Get("document")},
		LocalStorage:   WebStorage{storage: jsGlobal.Get("localStorage")},
		SessionStorage: WebStorage{storage: jsGlobal.Get("sessionStorage")},
		Path2D:         Path2D{path2D: jsGlobal.Get("Path2D")},
=======
		Document:       Htmldoc{Element: Element{el: js.Global.Get("document")}},
		LocalStorage:   WebStorage{storage: js.Global.Get("localStorage")},
		SessionStorage: WebStorage{storage: js.Global.Get("sessionStorage")},
		Path2D:         Path2D{path2D: js.Global.Get("Path2D")},
>>>>>>> Update APIs
	}
}

// Blur removes the focus from the window.
func (w *Window) Blur() {
	w.window.Call("blur")
}

// Focus sets the focus to the window.
func (w *Window) Focus() {
	w.window.Call("focus")
}

// Confirm shows an confirm dialog window with an OK and Cancel button.
func (w *Window) Confirm() bool {
	return w.window.Call("confirm").Bool()
}

// Print outputs the entire window to the browser's print dialog.
func (w *Window) Print() {
	w.window.Call("print")
}

func (w *Window) ScrollTo(xpos, ypos int) {
	w.window.Call("scrollTo", js.ValueOf(xpos), js.ValueOf(ypos))
}

func (w *Window) InnerHeight() int {
	return w.window.Get("innerHeight").Int()
}
