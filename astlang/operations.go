package astlang

import (
	"github.com/influx6/describe"
)

type CanName interface {
	SetName(name string)
}

func UseName(target *describe.Description, name string)   {
	if canName, ok := target.Get().(CanName); ok {
		canName.SetName(name)
		return
	}
	target.SetErr(describe.ErrNotApplicable)
}

type CanDescription interface {
	SetDescription(name string)
}

func UseDescription(target *describe.Description, desc string)   {
	if can, ok := target.Get().(CanDescription); ok {
		can.SetDescription(desc)
		return
	}
	target.SetErr(describe.ErrNotApplicable)
}

type CanVersion interface {
	SetVersion(name string)
}

func UseVersion(target *describe.Description, version string)   {
	if can, ok := target.Get().(CanVersion); ok {
		can.SetVersion(version)
		return
	}
	target.SetErr(describe.ErrNotApplicable)
}

func UseBaseType(target *describe.Description, baseType BaseType)   {
	if typeDefinition, ok := target.Get().(*TypeDefinition); ok {
		typeDefinition.Type = baseType
		return
	}
	target.SetErr(describe.ErrNotApplicable)
}

func UseMemory(target *describe.Description, mem MemoryLayout)   {
	if typeDefinition, ok := target.Get().(*TypeDefinition); ok {
		typeDefinition.Memory = mem
		return
	}
	target.SetErr(describe.ErrNotApplicable)
}

func UseAnnotation(target *describe.Description, text string, fn func()) AnnotationDefinition {
	var field AnnotationDefinition
	field.Content = text
	target.Push(&field)
	fn()
	return field
}

func UseComment(target *describe.Description, fn func()) CommentDefinition {
	var field CommentDefinition
	target.Push(&field)
	fn()
	return field
}

func UseCommentText(target *describe.Description, text string) {
	if comment, ok := target.Get().Elem().(*CommentDefinition); ok {
		comment.Contents = append(comment.Contents, text)
	}
}

func UseValue(target *describe.Description, value interface{})  {
	if obj, ok := target.Get().Elem().(*VariableDefinition); ok {
		obj.Value = value
	}
}

func UseConstant(target *describe.Description, fn func()) VariableDefinition {
	var field VariableDefinition
	field.Constant = true
	target.Push(&field)
	fn()
	return field
}

func UseVariable(target *describe.Description, fn func()) VariableDefinition {
	var field VariableDefinition
	target.Push(&field)
	fn()
	return field
}

func UseReturn(target *describe.Description, fn func()) ReturnDefinition {
	var field ReturnDefinition
	target.Push(&field)
	fn()
	return field
}

func UseResult(target *describe.Description, fn func()) ResultDefinition {
	var field ResultDefinition
	target.Push(&field)
	fn()
	return field
}

func UseField(target *describe.Description, fn func()) FieldDefinition {
	var field FieldDefinition
	target.Push(&field)
	fn()
	return field
}

func UseDataType(target *describe.Description, fn func()) DataTypeDefinition {
	var dt DataTypeDefinition
	target.Push(&dt)
	fn()
	return dt
}

func UseData(target *describe.Description, fn func()) DataDefinition {
	var dt DataDefinition
	target.Push(&dt)
	fn()
	return dt
}

func UseMethod(target *describe.Description, fn func()) MethodDefinition {
	var obj MethodDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseMethodCall(target *describe.Description, fn func()) MethodCallDefinition {
	var obj MethodCallDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseFor(target *describe.Description, fn func()) ForDefinition {
	var obj ForDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseLoop(target *describe.Description, fn func()) LoopDefinition {
	var obj LoopDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseIf(target *describe.Description, fn func()) IfDefinition {
	var obj IfDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseSwitch(target *describe.Description, fn func()) SwitchDefinition {
	var obj SwitchDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseCase(target *describe.Description, fn func()) CaseDefinition {
	var obj CaseDefinition
	target.Push(&obj)
	fn()
	return obj
}

type CanBody interface {
	SetBody(describe.Applicable)
}

func UseBody(target *describe.Description, body describe.Applicable)  {
	if canBody, ok := target.Get().Elem().(CanBody); ok {
		canBody.SetBody(body)
	}
}

func UseCondition(target *describe.Description, fn func()) ConditionDefinition {
	var obj ConditionDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseOperator(target *describe.Description, operator Operator, fn func()) OperatorDefinition {
	var obj OperatorDefinition
	obj.Operator = operator
	target.Push(&obj)
	fn()
	return obj
}
