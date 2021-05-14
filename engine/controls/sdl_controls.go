package controls

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLKeyboardControls struct {
	keys map[string]SDLKey
}

// NewSDLKeyboardController ...
func NewSDLKeyboardControls() *SDLKeyboardControls {
	return &SDLKeyboardControls{
		keys: map[string]SDLKey{},
	}
}

// SDLKey ...
type SDLKey struct {
	code       int
	isDown     bool
	justPushed bool
}

// Update ...
func (ctl *SDLKeyboardControls) Update() {
	sdlKeyboardState := sdl.GetKeyboardState()

	newKeys := map[string]SDLKey{}
	for name, oldKey := range ctl.keys {
		newKeys[name] = ctl.getNewKeyFromState(sdlKeyboardState, oldKey)
	}
	ctl.keys = newKeys
}

func (ctl SDLKeyboardControls) getNewKeyFromState(sdlKeyboardState []uint8, key SDLKey) SDLKey {
	return SDLKey{
		code:       key.code,
		isDown:     ctl.isKeyDownInSDL(sdlKeyboardState, key.code),
		justPushed: !key.isDown,
	}
}

func (ctl SDLKeyboardControls) isKeyDownInSDL(sdlState []uint8, keyCode int) bool {
	return sdlState[keyCode] == 1
}

// IsDown ...
func (ctl *SDLKeyboardControls) IsDown(keyName string) bool {
	key, ok := ctl.keys[keyName]
	if !ok {
		fmt.Println("No key with name:", keyName)
		return false
	}
	return key.isDown
}

// WasJustPushed ...
func (ctl *SDLKeyboardControls) WasJustPushed(keyName string) bool {
	key, ok := ctl.keys[keyName]
	if !ok {
		fmt.Println("No key with name:", keyName)
		return false
	}
	return key.justPushed
}

// RegisterKey ...
func (ctl *SDLKeyboardControls) RegisterKey(name string, code sdl.Keycode) {
	ctl.keys[name] = SDLKey{
		code:   int(code),
		isDown: false,
	}
}
