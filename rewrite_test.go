package rewrite_test

import (
	"errors"
	"testing"

	"github.com/influx6/rewrite"
	"github.com/stretchr/testify/require"
)

func TestStackDefinitions(t *testing.T) {
	var jsondef JSONDefinition
	var res, err = rewrite.Define(func(obj *rewrite.Description) {
		Object(obj, func() {
			Name(obj, "Nature")
			Desc(obj, "Desc")
		})
	})(&jsondef)

	require.NoError(t, err)
	require.NotNil(t, res)

	require.NotNil(t, jsondef.Target)
	require.Equal(t, jsondef.Target.Desc, "Desc")
	require.Equal(t, jsondef.Target.Name, "Nature")
}

type JSONDefinition struct {
	Target *Rob
}

func (j JSONDefinition) Elem() interface{} {
	return j
}

func (j *JSONDefinition) Apply(v interface{}) error {
	return nil
}

type Rob struct {
	Name string
	Desc string
}

func (j Rob) Elem() interface{} {
	return j
}

func (j *Rob) Apply(v interface{}) error {
	return nil
}

func Object(r *rewrite.Description, fn func()) {
	var obj Rob

	// Push object into task so functions
	// can apply them to it.
	r.Push(&obj)
	fn()
	// Pop object from stack, so we restore parent
	// context
	r.Pop()

	var parent, err = r.Get()
	if err != nil {
		return
	}

	objParent, ok := parent.(*JSONDefinition)
	if !ok {
		r.SetErr(errors.New("parent is not an object type"))
		return
	}

	objParent.Target = &obj
}

func Name(r *rewrite.Description, name string) {
	var parent, err = r.Get()
	if err != nil {
		return
	}

	objParent, ok := parent.(*Rob)
	if !ok {
		r.SetErr(errors.New("parent is not an object type"))
		return
	}

	objParent.Name = name
}

func Desc(r *rewrite.Description, desc string) {
	var parent, err = r.Get()
	if err != nil {
		return
	}

	objParent, ok := parent.(*Rob)
	if !ok {
		r.SetErr(errors.New("parent is not an object type"))
		return
	}

	objParent.Desc = desc
}
