package core

import "fmt"

var (
	ErrBehaviorMustHaveName = fmt.Errorf("name must not be empty")
	ErrBehaviorIsNotUnique  = fmt.Errorf("a behavior with this name is already attached")
	ErrBehaviorDoesNotExist = fmt.Errorf("no behavior with this name")

	ErrDuplicateEntity = fmt.Errorf("this entity is already a child")
	ErrNoSuchEntity    = fmt.Errorf("no child entity with that name")
)
