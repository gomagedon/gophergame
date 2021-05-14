package canvas

// Canvas ...
type Canvas interface {
	SetColor(r, g, b, a int)
	Clear()
	Present()
	DrawRect(x, y, w, h int)
}
