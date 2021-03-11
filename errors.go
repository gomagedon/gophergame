package engine

import "errors"

// ErrRendererMustHaveType ...
var ErrRendererMustHaveType = errors.New("This renderer does not have a type")

// ErrRendererMustBeUnique ...
var ErrRendererMustBeUnique = errors.New("A renderer of this type is already attached")

// ErrComponentMustHaveType ...
var ErrComponentMustHaveType = errors.New("This component does not have a type")

// ErrComponentMustBeUnique ...
var ErrComponentMustBeUnique = errors.New("A component of this type is already attached")
