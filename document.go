package browser

// Htmldoc is the root node of the HTML document.
type Htmldoc struct {
	Element
}

// GetElementByID receives an element from the DOM by its Id.
func (h Htmldoc) GetElementByID(id string) Element {
	return Element{
		el: h.el.Call("getElementById", id),
	}
}

// GetElementsByTagName receives an element from the DOM by its Id.
func (h Htmldoc) GetElementsByTagName(id string) Element {
	return Element{
		el: h.el.Call("getElementsByTagName", id),
	}
}

// To be implemented
func (Htmldoc) ActiveElement() {}

// func (Htmldoc) AddEventListener() {}
func (Htmldoc) AdoptNode() {}
func (Htmldoc) Anchors()   {}
func (Htmldoc) Applets()   {}
func (Htmldoc) BaseURI()   {}
func (Htmldoc) Body()      {}
func (h *Htmldoc) Close() {
	h.el.Call("close")
}
func (Htmldoc) Cookie()                 {}
func (Htmldoc) Charset()                {}
func (Htmldoc) CharacterSet()           {}
func (Htmldoc) CreateAttribute()        {}
func (Htmldoc) CreateComment()          {}
func (Htmldoc) CreateDocumentFragment() {}
func (Htmldoc) CreateEvent()            {}
func (Htmldoc) CreateTextNode()         {}
func (Htmldoc) DefaultView()            {}
func (Htmldoc) DesignMode()             {}
func (Htmldoc) Doctype()                {}
func (Htmldoc) DocumentElement()        {}
func (Htmldoc) DocumentMode()           {}
func (Htmldoc) DocumentURI()            {}
func (Htmldoc) Domain()                 {}
func (Htmldoc) DomConfig()              {}
func (Htmldoc) Embeds()                 {}
func (Htmldoc) ExecCommand()            {}
func (Htmldoc) Forms()                  {}
func (Htmldoc) GetElementsByClassName() {}
func (Htmldoc) GetElementsByName()      {}
func (Htmldoc) HasFocus()               {}
func (Htmldoc) Head()                   {}
func (Htmldoc) Images()                 {}
func (Htmldoc) Implementation()         {}
func (Htmldoc) ImportNode()             {}
func (Htmldoc) InputEncoding()          {}
func (Htmldoc) LastModified()           {}
func (Htmldoc) Links()                  {}
func (Htmldoc) Normalize()              {}
func (Htmldoc) NormalizeDocument()      {}
func (h *Htmldoc) Open() {

	h.el.Call("open")

}

// QuerySelector returns the first matching element where class = s.
func (h *Htmldoc) QuerySelector(s string) Element {

	return Element{el: h.el.Call("querySelector", s)}

}

// QuerySelectorAll returns the list of elements where class = s.
func (h *Htmldoc) QuerySelectorAll(s string) []Element {

	js := h.el.Call("querySelectorAll", s)

	e := make([]Element, js.Length())
	for i := 0; i < js.Length(); i++ {
		e[i] = Element{el: js.Index(i)}
	}

	return e
}

func (Htmldoc) ReadyState()          {}
func (Htmldoc) Referrer()            {}
func (Htmldoc) RemoveEventListener() {}
func (Htmldoc) RenameNode()          {}
func (Htmldoc) Scripts()             {}
func (Htmldoc) StrictErrorChecking() {}
func (Htmldoc) Title()               {}
func (Htmldoc) URL()                 {}
func (h *Htmldoc) Write(s string) {

	h.el.Call("write", s)

}
func (Htmldoc) Writeln() {}
