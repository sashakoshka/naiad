package naiad

type Style struct {
	thickness float64
	open      bool
}

func (style *Style) Thickness () (thickness float64) {
	return style.thickness
}

func (style *Style) Open () (closed bool) {
	return style.open
}

func (style *Style) SetThickness (thickness float64) {
	style.thickness = thickness
}

func (style *Style) SetOpen (open bool) {
	style.open = open
}
