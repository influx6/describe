# Describe
Describe provides a base library for constructing a stack-based definition package which allows creating functions chains
that are easily expressed to create complex structures.

The whole idea is one can define re-usable blocks of functions that define and augment a type, these created set can then be re-applied as many times as desirable to any instance of any type that supports it's expected interface contract. This means we can construct a DSL format which can be used to express different structures easily with the added benefit of organization and reuse. The contract then allows the receiving type to define what it expects, the rules for it's defining parts.

The goal of this package is to provide a base DSL which can be used to expressed language constructs for Models, Methods and Services, which then can be applied to a custom type which defines the rules for which such constructs can be applied and how that can be transformed into some other format (e.g a different language: javascript, rust, go, grahpql schemes). 

```go
var jsondef JSONDefinition
var res, err = describe.Define(func(obj *describe.DefStack) {
    Object(obj, func() {
        Name(obj, "Nature")
        Desc(obj, "Desc")
    })
})(&jsondef)
```

```go
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

func Object(r *describe.Description, fn func()) {
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

func Name(r *describe.Description, name string) {
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

func Desc(r *describe.Description, desc string) {
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
```
