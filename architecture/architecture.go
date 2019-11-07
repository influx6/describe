package architecture

import "github.com/influx6/restark"

// Name sets giving name attribute of pending
func Name(obj *restark.DefStack, value string) {
	obj.Apply(BaseName(value))
}

// BaseName applies giving name value to provided Applicable.
type BaseName string

// Version sets giving version attribute of reciever.
func Version(obj *restark.DefStack, value string) {
	obj.Apply(BaseVersion(value))
}

// Description sets giving description attribute of receiver.
func Description(obj *restark.DefStack, value string) {
	obj.Apply(BaseDescription(value))
}

// BaseDescription applies giving version value to provided Applicable.
type BaseDescription string

//****************************************
// MethodDefinitions
//****************************************

type BaseDefinition struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

// Apply applies giving value to base definition.
func (b *BaseDefinition) Apply(v interface{}) error {
	switch tm := v.(type) {
	case BaseName:
		b.Name = string(tm)
		return nil
	case BaseVersion:
		b.Version = string(tm)
		return nil
	case BaseDescription:
		b.Description = string(tm)
		return nil
	default:
		return nil
	}
}

// MethodDefinition defines the base definition for methods.
type MethodDefinition struct {
	BaseDefinition
}

// FieldDefinition defines the base definition for fields.
type FieldDefinition struct {
	BaseDefinition
}

// TypeDefinition defines the base definition for types.
type TypeDefinition struct {
	BaseDefinition
}

func (td TypeDefinition) Elem() interface{} {
	return td
}

func (td *TypeDefinition) Apply(item interface{}) error {
	return nil
}

// Type implements and adds a type into current stack with
// provided name, executing provided function.
func Type(obj *restark.DefStack, typeName string, fn func()) {
	defer obj.Release()

	var typeDef TypeDefinition
	typeDef.Name = typeName

	obj.Push(&typeDef)
	if fn != nil {
		fn()
	}
}
