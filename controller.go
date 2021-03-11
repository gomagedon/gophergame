package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var controller Controller

// Controller ...
type Controller interface {
	Update()
	IsDown(string) bool
	WasJustPushed(string) bool
}

// SDLKeyboardController ...
type SDLKeyboardController struct {
	keys map[string]SDLKey
}

// NewSDLKeyboardController ...
func NewSDLKeyboardController() *SDLKeyboardController {
	ctl := new(SDLKeyboardController)
	ctl.keys = map[string]SDLKey{}
	return ctl
}

// SDLKey ...
type SDLKey struct {
	code       sdl.Keycode
	isDown     bool
	justPushed bool
}

// Update ...
func (ctl *SDLKeyboardController) Update() {
	kbState := sdl.GetKeyboardState()

	for name, key := range ctl.keys {
		key.isDown = kbState[int(key.code)] == 1
		if key.isDown {
			key.justPushed = !ctl.keys[name].isDown
		}
		ctl.keys[name] = key
	}
}

// IsDown ...
func (ctl *SDLKeyboardController) IsDown(keyName string) bool {
	key, ok := ctl.keys[keyName]
	if !ok {
		fmt.Println("No key with name:", keyName)
		return false
	}
	return key.isDown
}

// WasJustPushed ...
func (ctl *SDLKeyboardController) WasJustPushed(keyName string) bool {
	key, ok := ctl.keys[keyName]
	if !ok {
		fmt.Println("No key with name:", keyName)
		return false
	}
	return key.justPushed
}

// RegisterKey ...
func (ctl *SDLKeyboardController) RegisterKey(name string, code sdl.Keycode) {
	ctl.keys[name] = SDLKey{
		code:   code,
		isDown: false,
	}
}
