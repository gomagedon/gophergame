package controls

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLKey struct {
	code       int
	isDown     bool
	justPushed bool
}

// SDLKeyboardControls is the SDL2 implementation of Controls
type SDLKeyboardControls struct {
	keys map[string]SDLKey
}

func NewSDLKeyboardControls() *SDLKeyboardControls {
	return &SDLKeyboardControls{
		keys: map[string]SDLKey{},
	}
}

/* Public Methods */

func (ctl SDLKeyboardControls) IsDown(keyName string) bool {
	key, ok := ctl.keys[keyName]
	if !ok {
		fmt.Println("No key with name:", keyName)
		return false
	}
	return key.isDown
}

func (ctl *SDLKeyboardControls) RegisterKey(name string, code sdl.Keycode) {
	ctl.keys[name] = SDLKey{
		code:   int(code),
		isDown: false,
	}
}

func (ctl *SDLKeyboardControls) Update() {
	sdlKeyboardState := sdl.GetKeyboardState()

	newKeys := map[string]SDLKey{}
	for name, oldKey := range ctl.keys {
		newKeys[name] = ctl.getNewKeyFromState(sdlKeyboardState, oldKey)
	}
	ctl.keys = newKeys
}

func (ctl SDLKeyboardControls) WasJustPushed(keyName string) bool {
	key, ok := ctl.keys[keyName]
	if !ok {
		fmt.Println("No key with name:", keyName)
		return false
	}
	return key.justPushed
}

/* Private Methods */

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
