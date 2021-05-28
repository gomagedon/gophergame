package engine

import "fmt"

var (
	ErrBehaviorMustHaveName       = fmt.Errorf("name must not be empty")
	ErrBehaviorMustHaveUniqueName = fmt.Errorf("a behavior with this name is already attached")
	ErrBehaviorDoesNotExist       = fmt.Errorf("no behavior with this name")

	ErrDuplicateEntity = fmt.Errorf("this entity is already a child")
	ErrNoSuchEntity    = fmt.Errorf("no child entity with that name")
)
