package generators

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/rewrite"
)

func Render(pkg rewrite.PackageDefinition) *jen.File {
	var code = jen.NewFile(pkg.GetName())
	code.Comment(pkg.GetDescription())
	code.Comment("\n")

	for _, definition := range pkg.Definitions {
		if err := render(code.Group, definition); err != nil {
			//TODO: Handle this better
			panic(err)
		}
	}
	return code
}

func render(group *jen.Group, definition rewrite.Applicable) error {
	switch def := definition.Elem().(type) {
	case rewrite.Value:
		return renderValue(group, def)
	case rewrite.VariableDefinition:
		return renderVariable(group, def)
	case rewrite.AssignmentDefinition:
		return renderAssignment(group, def)
	case rewrite.AnnotationDefinition:
		return renderAnnotationDefinition(group, def)
	case rewrite.CommentDefinition:
		return renderCommentDefinition(group, def)
	case rewrite.OperatorDefinition:
		return renderOperatorDefinition(group, def)
	case rewrite.DataDefinition:
		return renderDataDefinition(group, def)
	case rewrite.DataTypeDefinition:
		return renderDataTypeDefinition(group, def)
	case rewrite.ConditionDefinition:
		return renderConditionDefinition(group, def)
	case rewrite.IfDefinition:
		return renderIfDefinition(group, def)
	case rewrite.ReturnDefinition:
		return renderReturnDefinition(group, def)
	case rewrite.ForDefinition:
		return renderForDefinition(group, def)
	case rewrite.CaseDefinition:
		return renderCaseDefinition(group, def)
	case rewrite.SwitchDefinition:
		return renderSwitchDefinition(group, def)
	case rewrite.LiteralDefinition:
		return renderLiteralDefinition(group, def)
	case rewrite.FieldDefinition:
		return renderFieldDefinition(group, def)
	case rewrite.FutureDefinition:
		return renderFutureDefinition(group, def)
	case rewrite.ChannelDefinition:
		return renderChannelDefinition(group, def)
	case rewrite.StreamDefinition:
		return renderStreamDefinition(group, def)
	case rewrite.Function:
		return renderFunctionDefinition(group, def)
	case rewrite.AsyncFunction:
		return renderAsyncFunctionDefinition(group, def)
	default:
		return nerror.New("unable to handle %#v", definition)
	}
}

func renderAsyncFunctionDefinition(group *jen.Group, def rewrite.AsyncFunction) error {
	return nil
}

func renderFunctionDefinition(group *jen.Group, def rewrite.Function) error {
	group.Comment(def.Description).Func().Id(def.Name)

	if len(def.Accepts) != 0 {
		group.ParamsFunc(func(group *jen.Group) {
			for _, item := range def.Accepts {
				renderFuncArgumentDefinition(group, item)
			}
		})
	}

	if len(def.Returns) != 0 {
		group.ParamsFunc(func(group *jen.Group) {
			for _, item := range def.Returns {
				renderFunctionReturnDefinition(group, item)
			}
		})
	}

	group.BlockFunc(func(group *jen.Group) {
		for _, item := range def.Body {
			renderFuncBody(group, item)
		}
	})

	return nil
}

func renderFuncBody(group *jen.Group, def rewrite.Applicable) error {
	return nil
}

func renderFuncArgumentDefinition(group *jen.Group, def rewrite.VariableDefinition) error {
	return nil
}

func renderFunctionReturnDefinition(group *jen.Group, def rewrite.ResultDefinition) error {
	return nil
}

func renderFunctionTypeReturnDefinition(group *jen.Group, def rewrite.ResultDefinition) error {
	return nil
}

func renderReturnDefinition(group *jen.Group, def rewrite.ReturnDefinition) error {

	return nil
}

func renderOperatorDefinition(group *jen.Group, def rewrite.OperatorDefinition) error {
	renderOperator(group, def.Operator)
	return nil
}

func renderOperator(group *jen.Group, op rewrite.Operator) {
	switch op {
	case rewrite.ShortEqual:
		group.Op(":=")
		return
	case rewrite.Equal:
		group.Op("=")
		return
	case rewrite.Equality:
		group.Op("==")
		return
	case rewrite.NotEquality:
		group.Op("!=")
		return
	case rewrite.Increment:
		group.Op("++")
		return
	case rewrite.Decrement:
		group.Op("--")
		return
	case rewrite.Multiplication:
		group.Op("*")
		return
	case rewrite.Subtraction:
		group.Op("-")
		return
	case rewrite.Division:
		group.Op("/")
		return
	case rewrite.Addition:
		group.Op("+")
		return
	case rewrite.SelfMultiplication:
		group.Op("*=")
		return
	case rewrite.SelfSubtraction:
		group.Op("-=")
		return
	case rewrite.SelfDivision:
		group.Op("/=")
		return
	case rewrite.SelfAddition:
		group.Op("/=")
		return
	case rewrite.Modulo:
		group.Op("%")
		return
	case rewrite.LessThan:
		group.Op("<")
		return
	case rewrite.LessThanEqualTo:
		group.Op("<=")
		return
	case rewrite.GreaterThan:
		group.Op(">")
		return
	case rewrite.GreaterThanEqualTo:
		group.Op(">=")
		return
	case rewrite.ConditionalAnd:
		group.Op("&&")
		return
	case rewrite.ConditionalOR:
		group.Op("&&")
		return
	case rewrite.BinaryAnd:
		group.Op("&")
		return
	case rewrite.BinaryOR:
		group.Op("|")
		return
	case rewrite.BitwiseAnd:
		group.Op("&")
		return
	case rewrite.BitwiseNot:
		group.Op("~")
		return
	case rewrite.BitwiseOR:
		group.Op("|")
		return
	case rewrite.BitwiseXOR:
		group.Op("^")
		return
	case rewrite.LeftShift:
		group.Op("<<")
		return
	case rewrite.RightShift:
		group.Op(">>")
		return
	case rewrite.PointerAnd:
		group.Op("&")
		return
	case rewrite.PointerAsterick:
		group.Op("*")
		return
	}
}

func renderStreamDefinition(group *jen.Group, def rewrite.StreamDefinition) error {
	return nil
}

func renderChannelDefinition(group *jen.Group, def rewrite.ChannelDefinition) error {
	return nil
}

func renderFutureDefinition(group *jen.Group, def rewrite.FutureDefinition) error {
	return nil
}

func renderFieldDefinition(group *jen.Group, def rewrite.FieldDefinition) error {
	return nil
}

func renderSwitchDefinition(group *jen.Group, def rewrite.SwitchDefinition) error {

	return nil
}

func renderCaseDefinition(group *jen.Group, def rewrite.CaseDefinition) error {

	return nil
}

func renderForDefinition(group *jen.Group, def rewrite.ForDefinition) (err error) {
	group.ForFunc(func(group *jen.Group) {
		err = render(group, def.Start)
		if def.Middle != nil && err == nil {
			render(group, def.Middle)
		}
		if def.End != nil && err == nil {
			render(group, def.End)
		}
	})

	group.BlockFunc(func(group *jen.Group) {
		err = render(group, def.Body)
	})
	return
}

func renderIfDefinition(group *jen.Group, def rewrite.IfDefinition) (err error) {
	group.IfFunc(func(group *jen.Group) {
		renderConditionDefinition(group, def.Condition)
	})
	group.BlockFunc(func(group *jen.Group) {
		err = render(group, def.Body)
	})
	return
}

func renderConditionDefinition(group *jen.Group, def rewrite.ConditionDefinition) error {
	group.ParamsFunc(func(group *jen.Group) {
		render(group, def.Left)
		renderOperatorDefinition(group, def.Operator)
		render(group, def.Right)
	})
	return nil
}

func renderLiteralDefinition(group *jen.Group, def rewrite.LiteralDefinition) error {
	group.Comment(def.Description).Lit(def.Literal)
	return nil
}

// renders a typedef e.g type a int
func renderDataTypeDefinition(group *jen.Group, def rewrite.DataTypeDefinition) error {
	group.Type().Id(def.Name).Id(def.Type)
	return nil
}

// renders a struct
func renderDataDefinition(group *jen.Group, def rewrite.DataDefinition) error {
	group.Comment(def.Description).
		Type().Id(def.Name).
		StructFunc(func(fieldGroup *jen.Group) {
			for _, field := range def.Fields {
				render(group, field)
			}
		})

	for _, method := range def.Methods {
		renderDataMethodDefinition(group, def, method)
	}

	return nil
}

const dataObjectReceiverName = "dto"

func renderDataMethodReceiver(group *jen.Statement, data rewrite.DataDefinition, def rewrite.Function) *jen.Statement {
	return group.ParamsFunc(func(param *jen.Group) {
		if def.Pointer {
			param.Id(fmt.Sprintf("%s *%s", dataObjectReceiverName, data.Name))
			return
		}
		param.Id(fmt.Sprintf("%s %s", dataObjectReceiverName, data.Name))
	})
}

func renderDataMethodDefinition(group *jen.Group, data rewrite.DataDefinition, def rewrite.Function) error {
	renderDataMethodReceiver(group.Comment(def.Description).Func(), data, def).Id(def.Name)

	if len(def.Accepts) != 0 {
		group.ParamsFunc(func(group *jen.Group) {
			for _, item := range def.Accepts {
				renderFuncArgumentDefinition(group, item)
			}
		})
	}

	if len(def.Returns) != 0 {
		group.ParamsFunc(func(group *jen.Group) {
			for _, item := range def.Returns {
				renderFunctionReturnDefinition(group, item)
			}
		})
	}

	group.BlockFunc(func(group *jen.Group) {
		for _, item := range def.Body {
			renderFuncBody(group, item)
		}
	})

	return nil
}

func renderCommentDefinition(group *jen.Group, def rewrite.CommentDefinition) error {
	for _, content := range def.Contents {
		group.Comment(content)
	}
	return nil
}

func renderVariable(group *jen.Group, variableDefinition rewrite.VariableDefinition) error {
	group.Var().Id(variableDefinition.GetName())
	if err := render(group, variableDefinition.Type); err != nil {
		return err
	}

	if variableDefinition.Assign != nil {
		return render(group, variableDefinition.Assign)
	}

	return nil
}

func renderAssignment(group *jen.Group, def rewrite.AssignmentDefinition) error {
	if !def.Short {
		renderOperator(group, rewrite.Equal)
	}
	if def.Short {
		renderOperator(group, rewrite.ShortEqual)
	}
	return render(group, def.Value)
}

// renderValue renders the concrete value attached to the rewrite.Value object.
// For now it handles only go basic unit types (int, string, ...etc).
func renderValue(group *jen.Group, def rewrite.Value) error {
	// TODO: Consider adding custom handlers for non-default values.
	//switch realValue := def.Value.(type) {
	//}

	group.Lit(def.Value).Commentf("Value for %q", def.Name)
	return nil
}

func renderAnnotationDefinition(group *jen.Group, def rewrite.AnnotationDefinition) error {
	if len(def.Content) > 0 {
		group.Commentf("@%s(%s)", def.Name, def.Content)
		return nil
	}
	group.Commentf("@%s", def.Name)
	return nil
}
