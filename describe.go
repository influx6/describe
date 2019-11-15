package restark

import (
	"errors"

	"github.com/influx6/npkg/nerror"
)

// Applicable defines the base type which is to be applied to
// given function.
type Applicable interface {
	Apply(interface{}) error
}

// ReApplicable applies some operation to the Applicable.
type ReApplicable interface {
	Apply(Applicable) error
}

// DefStack manages a stack of Applicable implementing
// objects which allows popping and pushing value.
type DefStack struct {
	err    error
	stacks []Applicable
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

// Push adds a new item into the Applicable list.
func (s *DefStack) Push(item interface{}) {
	s.stacks = append(s.stacks, item)
}

// First returns first Applicable object in stack.
// Usually the first Applicable is the source and
// root of all defined Definitions.
func (s *DefStack) First() (Applicable, error) {
	if len(s.stacks) == 0 {
		return nil, nerror.New("Empty Applicable")
	}
	return s.stacks[0], nil
}

// Get returns current Applicable object in stack.
func (s *DefStack) Get() (Applicable, error) {
	if len(s.stacks) == 0 {
		return nil, nerror.New("Empty Applicable")
	}
	return s.stacks[len(s.stacks)-1], nil
}

// Pop pops recent stack to the last used stack.
// If called iteratively then all items will be removed from stack.
func (s *DefStack) Pop() Applicable {
	if len(s.stacks) == 0 {
		return nil
	}

	elem := s.stacks[len(s.stacks)-1]
	s.stacks = s.stacks[:len(s.stacks)-1]
	return elem
}

// Apply applies provided target to the provided current
// top element on the stack.
func (s *DefStack) Apply(target interface{}) error {
	var appl, err = s.Get()
	if err != nil {
		return err
	}

	if re, ok := target.(ReApplicable); ok {
		return re.Apply(appl)
	}

	return appl.Apply(target)
}

// Release will pop the current top elements on the stack
// applying it to it's parent.
func (s *DefStack) Release() {
	if len(s.stacks) == 1 || len(s.stacks) == 0 {
		return
	}

	var current = s.Pop()
	var parent, err = s.Get()
	if err != nil {
		s.SetErr(err)
		return
	}

	var applicable, ok = parent.(Applicable)
	if !ok {
		s.SetErr(nerror.New("parent is not an Applicable"))
		return
	}

	if err = parent.Apply(current); err != nil {
		s.SetErr(err)
	}
}

//**************************
// Applicable Functions
//**************************

// Definition defines a function type which represents
// an operation to be applied to a provided stack element.
type Definition func(stack *DefStack)

// DefinitionApplication defines the function type which represents the
// target for a definition function
type DefinitionApplication func(source interface{}) (interface{}, error)

// Compose combines giving series of definitions as a single Applicable.
//
// It returns a single function which can be re-used elsewhere.
func Compose(fns ...Definition) Definition {
	return func(root *DefStack) {
		for _, fn := range fns {
			fn(root)
		}
	}
}

// Pop returns a Definition which pops the last value from the
// stack.
func Pop() Definition {
	return func(root *DefStack) {
		root.Pop()
	}
}

// Push returns a Definition function which will always push
// given type into underline DefStack to be applied to the parent.
func Push(t Applicable) Definition {
	return func(root *DefStack) {
		root.Push(t)
	}
}

// ApplyToFirst returns a Definition which pops the last value on the stack
// applying it to root or first item within the stack.
func ApplyToFirst() Definition {
	return func(root *DefStack) {
		var last = root.Pop()
		var parent, err = root.First()
		if err != nil {
			root.SetErr(errors.New("no parent to apply Definition to"))
			return
		}

		if err = parent.Apply(last); err != nil {
			root.SetErr(err)
		}
	}
}

// ApplyToLast returns a Definition which pops the last value from the stack
// applying it to previous item within the stack.
func ApplyToLast() Definition {
	return func(root *DefStack) {
		var last = root.Pop()
		var parent, err = root.Get()
		if err != nil {
			root.SetErr(errors.New("no parent to apply Definition to"))
			return
		}

		if err = parent.Apply(last); err != nil {
			root.SetErr(err)
		}
	}
}

// Define returns a function matching the DefinitionApplication type.
// It applies all provided definitions to the source matching the Applicable interface
// which will host and validate definitions being applied to it.
//
// An error is returned if any definition (that is either the main or internals) fail
// to meet defined criteria.
func Define(fn ...Definition) DefinitionApplication {
	return func(source interface{}) (interface{}, error) {
		// Create the Applicable stack which will hold
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
