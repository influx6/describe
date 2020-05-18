package generators

import (
	"github.com/influx6/describe"
	"github.com/influx6/describe/astlang"
)
import "github.com/dave/jennifer/jen"

func Render(pkg astlang.PackageDefinition) *jen.File {
	var code = jen.NewFile(pkg.GetName())
	code.Comment(pkg.GetDescription())
	code.Comment("\n")
	code.Comment("Version: ")
	code.Comment(pkg.Version)

	for _, definition := range pkg.Definitions {
		render(code, definition)
	}
	return code
}

func render(file *jen.File, definition describe.Applicable) {
	switch def := definition.Elem().(type) {
	case astlang.VariableDefinition:
		renderVariable(file, def)
	case astlang.AssignmentDefinition:
		renderAssignment(file, def)
	case astlang.AnnotationDefinition:
	case astlang.CommentDefinition:
	case astlang.ResultDefinition:
	case astlang.DataDefinition:
	case astlang.DataTypeDefinition:
	case astlang.ConditionDefinition:
	case astlang.IfDefinition:
	case astlang.ReturnDefinition:
	case astlang.LoopDefinition:
	case astlang.ForDefinition:
	case astlang.CaseDefinition:
	case astlang.SwitchDefinition:
	case astlang.FieldDefinition:
	case astlang.FutureDefinition:
	case astlang.ChannelDefinition:
	case astlang.StreamDefinition:
	}
}

func renderVariable(file *jen.File, variableDefinition astlang.VariableDefinition) {
	file.Var().Id(variableDefinition.GetName())
	render(file, variableDefinition.Type)
	if variableDefinition.Assign != nil {
		render(file, variableDefinition.Assign)
	}
}

func renderAssignment(file *jen.File, def astlang.AssignmentDefinition) {

}