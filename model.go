package restark

// Name sets giving name attribute of pending
func Name(obj *DefStack, value string) {
	panic(obj.Apply(BaseName(value)))
}

// BaseName applies giving name value to provided Applicable.
type BaseName string

// Version sets giving version attribute of reciever.
func Version(obj *DefStack, value string) {
	panic(obj.Apply(BaseVersion(value)))
}

type BaseVersion string

// Description sets giving description attribute of receiver.
func Description(obj *DefStack, value string) {
	panic(obj.Apply(BaseDescription(value)))
}

// BaseDescription applies giving version value to provided Applicable.
type BaseDescription string

//****************************************
// MethodDefinitions
//****************************************

// BaseDefinition describes the base definition used by other
// packages.
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

// TypeDefinition defines the base definition for types.
type TypeDefinition struct {
	BaseDefinition

	Fields  []FieldDefinition
	Methods []MethodDefinition
}

func (t *TypeDefinition) Apply(elem interface{}) error {
	switch tm := elem.(type) {
	case FieldDefinition:
		t.Fields = append(t.Fields, tm)
	case MethodDefinition:
		t.Methods = append(t.Methods, tm)
	}
	return t.BaseDefinition.Apply(elem)
}

// Type implements and adds a type into current stack with
// provided name, executing provided function.
func Type(obj *DefStack, typeName string, fn func()) {
	defer obj.Release()

	var typeDef TypeDefinition
	typeDef.Name = typeName

	obj.Push(&typeDef)
	if fn != nil {
		fn()
	}
}

// FieldDefinition defines the base definition for fields.
type FieldDefinition struct {
	BaseDefinition

	Type TypeDefinition `json:"type"`
}

func (t *FieldDefinition) Apply(elem interface{}) error {
	switch tm := elem.(type) {
	case TypeDefinition:
		t.Type = tm
	}
	return t.BaseDefinition.Apply(elem)
}

// Field adds FieldDefinition to a giving stack.
func Field(obj *DefStack, fieldName string, fn func()) {
	defer obj.Release()

	var typeDef FieldDefinition
	typeDef.Name = fieldName

	obj.Push(&typeDef)
	if fn != nil {
		fn()
	}
}

// ArgumentDefinition defines the base definition for fields.
type ArgumentDefinition struct {
	BaseDefinition

	Type TypeDefinition `json:"type"`
}

func (t *ArgumentDefinition) Apply(elem interface{}) error {
	switch tm := elem.(type) {
	case TypeDefinition:
		t.Type = tm
	}
	return t.BaseDefinition.Apply(elem)
}

// ReturnDefinition defines the base definition for fields.
type ReturnDefinition struct {
	BaseDefinition

	Type TypeDefinition `json:"type"`
}

func (t *ReturnDefinition) Apply(elem interface{}) error {
	switch tm := elem.(type) {
	case TypeDefinition:
		t.Type = tm
	}
	return t.BaseDefinition.Apply(elem)
}

// MethodDefinition defines the base definition for methods.
type MethodDefinition struct {
	BaseDefinition

	Returns   []ReturnDefinition
	Arguments []ArgumentDefinition
}

func (t *MethodDefinition) Apply(elem interface{}) error {
	switch tm := elem.(type) {
	case ReturnDefinition:
		t.Returns = append(t.Returns, tm)
	case ArgumentDefinition:
		t.Arguments = append(t.Arguments, tm)
	}
	return t.BaseDefinition.Apply(elem)
}
