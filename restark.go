package restark

import (
	"github.com/influx6/npkg/nerror"
)

// Stacked defines the base type which is to be applied to
// given function.
type Stacked interface {
	Elem() interface{}

	Apply(interface{}) error
}

// StackApp manages a stack of Stacked implementing
// objects which allows popping and pushing value.
type StackApp struct {
	stacks []Stacked
}

// Apply applies provided value to the current top element
// within provided stack.
func (s *StackApp) Apply(v interface{}) error {
	var current, err = s.Peek()
	if err != nil {
		return err
	}
	return current.Apply(v)
}

// Elem returns the StackApp which ensures it matches the stacked
// interface.
func (s *StackApp) Elem() interface{} {
	return s
}

// Push adds a new item into the stacked list.
func (s *StackApp) Push(item Stacked) {
	s.stacks = append(s.stacks, item)
}

// Peek returns current Stacked object in stack.
func (s *StackApp) Peek() (Stacked, error) {
	if len(s.stacks) == 0 {
		return nil, nerror.New("Empty stacked")
	}
	return s.stacks[len(s.stacks)-1], nil
}

// Pop pops recent stack to the last used stack.
// If called iteratively then all items will be removed from stack.
func (s *StackApp) Pop() (Stacked, error) {
	if len(s.stacks) == 0 {
		return nil, nerror.New("Empty stacked")
	}

	elem := s.stacks[len(s.stacks)-1]
	s.stacks = s.stacks[:len(s.stacks)-1]
	return elem, nil
}
