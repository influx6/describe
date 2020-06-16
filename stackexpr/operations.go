package stackexpr

import (
	"github.com/influx6/rewrite"
)

type CanName interface {
	SetName(name string)
}

func UseName(target *Description, name string)   {
	if canName, ok := target.Get().(CanName); ok {
		canName.SetName(name)
		return
	}
	target.SetErr(rewrite.ErrNotApplicable)
}

type CanDescription interface {
	SetDescription(name string)
}

func UseDescription(target *Description, desc string)   {
	if can, ok := target.Get().(CanDescription); ok {
		can.SetDescription(desc)
		return
	}
	target.SetErr(rewrite.ErrNotApplicable)
}

type CanVersion interface {
	SetVersion(name string)
}

func UseVersion(target *Description, version string)   {
	if can, ok := target.Get().(CanVersion); ok {
		can.SetVersion(version)
		return
	}
	target.SetErr(rewrite.ErrNotApplicable)
}

func UseBaseType(target *Description, baseType rewrite.BaseType)   {
	if typeDefinition, ok := target.Get().(*rewrite.TypeDefinition); ok {
		typeDefinition.Type = baseType
		return
	}
	target.SetErr(rewrite.ErrNotApplicable)
}

func UseMemory(target *Description, mem rewrite.MemoryLayout)   {
	if typeDefinition, ok := target.Get().(*rewrite.TypeDefinition); ok {
		typeDefinition.Memory = mem
		return
	}
	target.SetErr(rewrite.ErrNotApplicable)
}

func UseAnnotation(target *Description, text string, fn func()) rewrite.AnnotationDefinition {
	var field rewrite.AnnotationDefinition
	field.Content = text
	target.Push(&field)
	fn()
	return field
}

func UseComment(target *Description, fn func()) rewrite.CommentDefinition {
	var field rewrite.CommentDefinition
	target.Push(&field)
	fn()
	return field
}

func UseCommentText(target *Description, text string) {
	if comment, ok := target.Get().Elem().(*rewrite.CommentDefinition); ok {
		comment.Contents = append(comment.Contents, text)
	}
}

func UseValue(target *Description, value interface{})  {
	//if obj, ok := target.Get().Elem().(*astlang.VariableDefinition); ok {
	//	//obj.Value = value
	//}
}

func UseConstant(target *Description, fn func()) rewrite.VariableDefinition {
	var field rewrite.VariableDefinition
	field.Constant = true
	target.Push(&field)
	fn()
	return field
}

func UseVariable(target *Description, fn func()) rewrite.VariableDefinition {
	var field rewrite.VariableDefinition
	target.Push(&field)
	fn()
	return field
}

func UseReturn(target *Description, fn func()) rewrite.ReturnDefinition {
	var field rewrite.ReturnDefinition
	target.Push(&field)
	fn()
	return field
}

func UseResult(target *Description, fn func()) rewrite.ResultDefinition {
	var field rewrite.ResultDefinition
	target.Push(&field)
	fn()
	return field
}

func UseField(target *Description, fn func()) rewrite.FieldDefinition {
	var field rewrite.FieldDefinition
	target.Push(&field)
	fn()
	return field
}

func UseDataType(target *Description, fn func()) rewrite.DataTypeDefinition {
	var dt rewrite.DataTypeDefinition
	target.Push(&dt)
	fn()
	return dt
}

func UseData(target *Description, fn func()) rewrite.DataDefinition {
	var dt rewrite.DataDefinition
	target.Push(&dt)
	fn()
	return dt
}

func UseMethod(target *Description, fn func()) rewrite.MethodDefinition {
	var obj rewrite.MethodDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseMethodCall(target *Description, fn func()) rewrite.MethodCallDefinition {
	var obj rewrite.MethodCallDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseFor(target *Description, fn func()) rewrite.ForDefinition {
	var obj rewrite.ForDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseLoop(target *Description, fn func()) rewrite.LoopDefinition {
	var obj rewrite.LoopDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseIf(target *Description, fn func()) rewrite.IfDefinition {
	var obj rewrite.IfDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseSwitch(target *Description, fn func()) rewrite.SwitchDefinition {
	var obj rewrite.SwitchDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseCase(target *Description, fn func()) rewrite.CaseDefinition {
	var obj rewrite.CaseDefinition
	target.Push(&obj)
	fn()
	return obj
}

type CanBody interface {
	SetBody(rewrite.Applicable)
}

func UseBody(target *Description, body rewrite.Applicable)  {
	if canBody, ok := target.Get().Elem().(CanBody); ok {
		canBody.SetBody(body)
	}
}

func UseCondition(target *Description, fn func()) rewrite.ConditionDefinition {
	var obj rewrite.ConditionDefinition
	target.Push(&obj)
	fn()
	return obj
}

func UseOperator(target *Description, operator rewrite.Operator, fn func()) rewrite.OperatorDefinition {
	var obj rewrite.OperatorDefinition
	obj.Operator = operator
	target.Push(&obj)
	fn()
	return obj
}
