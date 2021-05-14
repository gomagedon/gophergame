package controls

type Controls interface {
	Update()
	IsDown(string) bool
	WasJustPushed(string) bool
}
