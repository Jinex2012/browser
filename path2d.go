package browser

import (
	"syscall/js"
)

type Path2D struct {
	path2D js.Value
}

func (p Path2D) New(args ...interface{}) *Path2D {
	return &Path2D{path2D: p.path2D.New()}
}

// Experimental
func (p *Path2D) AddPath() {
	p.path2D.Call("addPath")
}

func (p *Path2D) ClosePath() {
	p.path2D.Call("closePath")
}

func (p *Path2D) MoveTo(x, y float64) {
	p.path2D.Call("moveTo", js.ValueOf(x), js.ValueOf(y))
}

func (p *Path2D) LineTo(x, y float64) {
	p.path2D.Call("lineTo", js.ValueOf(x), js.ValueOf(y))
}

func (p *Path2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64) {
	p.path2D.Call("bezierCurveTo",
		js.ValueOf(cp1x),
		js.ValueOf(cp1y),
		js.ValueOf(cp2x),
		js.ValueOf(cp2y),
		js.ValueOf(x),
		js.ValueOf(y))
}

func (p *Path2D) Arc(x, y, radius, startAngle, endAngle float64, counterclockwise bool) {
	p.path2D.Call("arc",
		js.ValueOf(x),
		js.ValueOf(y),
		js.ValueOf(radius),
		js.ValueOf(startAngle),
		js.ValueOf(endAngle),
		js.ValueOf(counterclockwise))

}

func (p *Path2D) ArcTo(x1, y1, x2, y2, r float64) {
	p.path2D.Call("arcTo", js.ValueOf(x1), js.ValueOf(y1), js.ValueOf(x2), js.ValueOf(y2), js.ValueOf(r))
}

func (p *Path2D) Ellipse() {
	p.path2D.Call("ellipse")
}

func (p *Path2D) Rect(x, y, w, h float64) {
	p.path2D.Call("rect", js.ValueOf(x), js.ValueOf(y), js.ValueOf(w), js.ValueOf(h))

}

func (p *Path2D) GetPath() js.Value {
	return p.path2D
}
