package entity

import "fmt"

var ErrBehaviorMustHaveType = fmt.Errorf("this behavior does not have a type")
var ErrBehaviorMustBeUnique = fmt.Errorf("a behavior of this type is already attached")
var ErrBehaviorDoesNotExist = fmt.Errorf("no behavior with this name")

var ErrDuplicateEntity = fmt.Errorf("this entity is already a child")
var ErrNoSuchEntity = fmt.Errorf("no child entity with that name")
