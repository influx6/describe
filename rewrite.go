package rewrite

import (
	"errors"
)

var ErrNotApplicable = errors.New("cant apply to target")

// DefinitionMiddleware defines the function type which represents the
// target for a definition function
type DefinitionMiddleware func(source Applicable) (Applicable, error)


// Applicable defines the base type which is to be applied to
// given function.
type Applicable interface {
	Elem() interface{}
}

// To allow stack-like operations on a series of Applicable.
type Stack interface {
	Release() (parent, child Applicable)
	Current() Applicable
	Root() Applicable
	Pop() Applicable
	Push(Applicable)
	IsUsable() bool
	SetErr(err error)
	Err() error
}

// Definition defines a function type which represents
// an operation to be applied to a provided stack element.
type Definition func(stack Stack)

