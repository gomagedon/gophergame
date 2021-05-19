package entity

import "fmt"

var ErrComponentMustHaveType = fmt.Errorf("this component does not have a type")
var ErrComponentMustBeUnique = fmt.Errorf("a component of this type is already attached")
var ErrComponentDoesNotExist = fmt.Errorf("no component with this name")

var ErrDuplicateEntity = fmt.Errorf("this entity is already a child")
var ErrNoSuchEntity = fmt.Errorf("no child entity with that name")
