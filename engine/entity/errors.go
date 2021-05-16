package entity

import "fmt"

var ErrRendererMustHaveType = fmt.Errorf("this renderer does not have a type")
var ErrRendererMustBeUnique = fmt.Errorf("a Errorfnderer of this type is already attached")
var ErrComponentMustHaveType = fmt.Errorf("this component does not have a type")
var ErrComponentMustBeUnique = fmt.Errorf("a component of this type is already attached")

var ErrDuplicateEntity = fmt.Errorf("this entity is already a child")
var ErrNoSuchEntity = fmt.Errorf("no child entity with that name")
