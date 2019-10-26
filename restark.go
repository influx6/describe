package restark

import (
	"github.com/influx6/npkg/nerror"
)

// Definition defines the base type which is to be applied to
// given function.
type Definition interface {
	Elem() interface{}
	Apply(interface{}) error
}

// DefStack manages a stack of Definition implementing
// objects which allows popping and pushing value.
type DefStack struct {
	err    error
	stacks []Definition
}

// Apply applies provided value to the current top element
// within provided stack.
func (s *DefStack) Apply(v interface{}) error {
	if s.err != nil {
		return s.err
	}

	var current, err = s.Get()
	if err != nil {
		return err
	}

	return current.Apply(v)
}

// SetErr sets the error returned when DefStack.Apply is
// called.
func (s *DefStack) SetErr(err error) {
	s.err = err
}

// Err returns possible attached error of a giving DefStack.
func (s *DefStack) Err() error {
	return s.err
}

// Elem returns the DefStack which ensures it matches the Definition
// interface.
func (s *DefStack) Elem() interface{} {
	return s
}

// Push adds a new item into the Definition list.
func (s *DefStack) Push(item Definition) {
	s.stacks = append(s.stacks, item)
}

// Get returns current Definition object in stack.
func (s *DefStack) Get() (Definition, error) {
	if len(s.stacks) == 0 {
		return nil, nerror.New("Empty Definition")
	}
	return s.stacks[len(s.stacks)-1], nil
}

// Pop pops recent stack to the last used stack.
// If called iteratively then all items will be removed from stack.
func (s *DefStack) Pop() Definition {
	if len(s.stacks) == 0 {
		return nil
	}

	elem := s.stacks[len(s.stacks)-1]
	s.stacks = s.stacks[:len(s.stacks)-1]
	return elem
}

//**************************
// Definition Functions
//**************************

// Definitions defines a function type which represents
// an operation to be applied to a provided stack element.
type Definitions func(stack *DefStack)

// DefinitionApplication defines the function type which represents the
// target for a definition function
type DefinitionApplication func(source Definition) (Definition, error)

// Define returns a function matching the DefinitionApplication type.
// It applies all provided definitions to the source matching the Definition interface
// which will host and validate definitions being applied to it.
//
// An error is returned if any definition (that is either the main or internals) fail
// to meet defined criteria.
func Define(fn ...Definitions) DefinitionApplication {
	return func(source Definition) (Definition, error) {
		// Create the Definition stack which will hold
		// ever increasing definitions from child to parent contexts.
		var stack DefStack
		stack.Push(source)

		// Apply definitions to stack.
		for _, fn := range fn {
			if stack.err != nil {
				return source, stack.err
			}

			fn(&stack)
		}

		if stack.err != nil {
			return source, stack.err
		}

		return source, nil
	}
}
