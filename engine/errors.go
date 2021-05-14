package engine

import "errors"

// ErrRendererMustHaveType ...
var ErrRendererMustHaveType = errors.New("this renderer does not have a type")

// ErrRendererMustBeUnique ...
var ErrRendererMustBeUnique = errors.New("a renderer of this type is already attached")

// ErrComponentMustHaveType ...
var ErrComponentMustHaveType = errors.New("this component does not have a type")

// ErrComponentMustBeUnique ...
var ErrComponentMustBeUnique = errors.New("a component of this type is already attached")
