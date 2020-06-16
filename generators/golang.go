package generators

import (
	"github.com/influx6/rewrite"
)
import "github.com/dave/jennifer/jen"

func Render(pkg rewrite.PackageDefinition) *jen.File {
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

func render(file *jen.File, definition rewrite.Applicable) {
	switch def := definition.Elem().(type) {
	case rewrite.VariableDefinition:
		renderVariable(file, def)
	case rewrite.AssignmentDefinition:
		renderAssignment(file, def)
	case rewrite.AnnotationDefinition:
	case rewrite.CommentDefinition:
	case rewrite.ResultDefinition:
	case rewrite.DataDefinition:
	case rewrite.DataTypeDefinition:
	case rewrite.ConditionDefinition:
	case rewrite.IfDefinition:
	case rewrite.ReturnDefinition:
	case rewrite.LoopDefinition:
	case rewrite.ForDefinition:
	case rewrite.CaseDefinition:
	case rewrite.SwitchDefinition:
	case rewrite.FieldDefinition:
	case rewrite.FutureDefinition:
	case rewrite.ChannelDefinition:
	case rewrite.StreamDefinition:
	}
}

func renderVariable(file *jen.File, variableDefinition rewrite.VariableDefinition) {
	file.Var().Id(variableDefinition.GetName())
	render(file, variableDefinition.Type)
	if variableDefinition.Assign != nil {
		render(file, variableDefinition.Assign)
	}
}

func renderAssignment(file *jen.File, def rewrite.AssignmentDefinition) {

}
