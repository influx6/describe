package architecture

import "github.com/influx6/restark"

//****************************************
// MethodDefinitions
//****************************************

// MethodDefinition defines the base definition for methods.
type MethodDefinition struct {
	BaseDefinition
}

//****************************************
// FieldDefinition
//****************************************

// FieldDefinition defines the base definition for fields.
type FieldDefinition struct {
	BaseDefinition
}

//****************************************
// TypeDefinition
//****************************************

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
