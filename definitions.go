package rewrite

type Meta struct {
	Version     string `json:"version"`
}

type Root struct {
	Meta        Meta
	Definitions []Applicable
}

type BaseDefinition struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

func (bd BaseDefinition) Elem() interface{} {
	return bd
}

func (bd *BaseDefinition) GetName() string {
	return bd.Name
}

func (bd *BaseDefinition) GetDescription() string {
	return bd.Description
}

func (bd *BaseDefinition) SetDescription(desc string)  {
	bd.Description = desc
}

func (bd *BaseDefinition) SetName(name string)  {
	bd.Name = name
}

func (bd *BaseDefinition) SetVersion(version string)  {
	bd.Version = version
}

func (bd *BaseDefinition) Apply(item interface{}) error {
	return ErrNotApplicable
}

const (
	Bit64 MemoryLayout = iota
	Bit32
)

type MemoryLayout int

const (
	Rune BaseType = iota + 1
	String
	Decimal
	Integer
	Complex
	Time
)

type BaseType int

func (b BaseType) String() string {
	switch b {
	case String:
		return "string"
	case Rune:
		return "rune"
	case Decimal:
		return "decimal"
	case Integer:
		return "integer"
	case Complex:
		return "complex"
	case Time:
		return "time"
	default:
		return "runtime"
	}
}

type PackageDefinition struct {
	BaseDefinition
	Definitions []Applicable
}

func (td PackageDefinition) Elem() interface{} {
	return td
}

func (td *PackageDefinition) Apply(item interface{}) error {
	if def, ok := item.(Applicable); ok {
		td.Definitions = append(td.Definitions, def)
	}
	return ErrNotApplicable
}

// TypeDefinition defines the base definition for types.
type TypeDefinition struct {
	BaseDefinition
	Type   BaseType
	Memory MemoryLayout
}

func (td TypeDefinition) Elem() interface{} {
	return td
}

func (td *TypeDefinition) Apply(item interface{}) error {
	switch ritem := item.(type) {
	case BaseType:
		td.Type = ritem
	case MemoryLayout:
		td.Memory = ritem
	case *BaseDefinition:
		td.BaseDefinition = *ritem
	case BaseDefinition:
		td.BaseDefinition = ritem
	}
	return ErrNotApplicable
}

// MethodDefinition defines the base definition for methods.
type MethodDefinition struct {
	BaseDefinition
	Arguments []FieldDefinition
	Returns   []ReturnDefinition
	Data      Applicable
}

func (td MethodDefinition) Elem() interface{} {
	return td
}

func (td *MethodDefinition) Apply(item interface{}) error {
	switch ritem := item.(type) {
	case *ReturnDefinition:
		td.Returns = append(td.Returns, *ritem)
		return nil
	case *FieldDefinition:
		td.Arguments = append(td.Arguments, *ritem)
		return nil
	case ReturnDefinition:
		td.Returns = append(td.Returns, ritem)
		return nil
	case FieldDefinition:
		td.Arguments = append(td.Arguments, ritem)
		return nil
	case *BaseDefinition:
		td.BaseDefinition = *ritem
		return nil
	case BaseDefinition:
		td.BaseDefinition = ritem
		return nil
	case Applicable:
		td.Data = ritem
	}
	return ErrNotApplicable
}

type MethodCallDefinition struct {
	BaseDefinition
	Arguments []VariableDefinition
	Results []ResultDefinition
	Method *MethodDefinition
}

func (td MethodCallDefinition) Elem() interface{} {
	return td
}

func (td *MethodCallDefinition) Apply(item interface{}) error {
	switch ritem := item.(type) {
	case *ResultDefinition:
		td.Results = append(td.Results, *ritem)
		return nil
	case ResultDefinition:
		td.Results = append(td.Results, ritem)
		return nil
	case *VariableDefinition:
		td.Arguments = append(td.Arguments, *ritem)
		return nil
	case VariableDefinition:
		td.Arguments = append(td.Arguments, ritem)
		return nil
	case *BaseDefinition:
		td.BaseDefinition = *ritem
		return nil
	case BaseDefinition:
		td.BaseDefinition = ritem
		return nil
	case *MethodDefinition:
		td.Method = ritem
		return nil
	case MethodDefinition:
		td.Method = &ritem
		return nil
	}
	return ErrNotApplicable
}

type CommentDefinition struct {
	BaseDefinition
	Contents []string
}

func (td CommentDefinition) Elem() interface{} {
	return td
}

func (td *CommentDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case string:
		td.Contents = append(td.Contents, value)
		return nil
	case []string:
		td.Contents = append(td.Contents, value...)
		return nil
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type AnnotationDefinition struct {
	BaseDefinition
	Content string
}

func (td AnnotationDefinition) Elem() interface{} {
	return td
}

func (td *AnnotationDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case string:
		td.Content = value
		return nil
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type ResultDefinition struct {
	BaseDefinition
	Type  Applicable
	Value Applicable
}

func (td ResultDefinition) Elem() interface{} {
	return td
}

func (td *ResultDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type Value struct {
	BaseDefinition
	Value Applicable
	Short bool // should use := instead of =
}

func (td Value) Elem() interface{} {
	return td
}

func (td *Value) Apply(item interface{}) error {
	return ErrNotApplicable
}

type AssignmentDefinition struct {
	BaseDefinition
	Value Applicable
	Short bool // should use := instead of =
}

func (td AssignmentDefinition) Elem() interface{} {
	return td
}

func (td *AssignmentDefinition) Apply(item interface{}) error {
	return ErrNotApplicable
}

type VariableDefinition struct {
	BaseDefinition
	Type     Applicable
	Assign   *AssignmentDefinition
	Constant bool
}

func (td VariableDefinition) Elem() interface{} {
	return td
}

func (td *VariableDefinition) Apply(item interface{}) error {
	return ErrNotApplicable
}

type ReturnDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td ReturnDefinition) Elem() interface{} {
	return td
}

func (td *ReturnDefinition) Apply(item interface{}) error {
	if value, ok := item.(Applicable); ok {
		td.Type = value
	}
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type FieldDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td FieldDefinition) Elem() interface{} {
	return td
}

func (td *FieldDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case Applicable:
		td.Type = value
		return nil
	}
	return ErrNotApplicable
}

// DataType represents a defined type where the type is a previously
// created data type.
type DataTypeDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td DataTypeDefinition) Elem() interface{} {
	return td
}

func (td *DataTypeDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case Applicable:
		td.Type = value
		return nil
	}
	return ErrNotApplicable
}

type DataDefinition struct {
	BaseDefinition
	Fields []FieldDefinition
	Methods []MethodDefinition
}

func (td DataDefinition) Elem() interface{} {
	return td
}

func (td *DataDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case []MethodDefinition:
		td.Methods = append(td.Methods, value...)
		return nil
	case []FieldDefinition:
		td.Fields = append(td.Fields, value...)
		return nil
	case MethodDefinition:
		td.Methods = append(td.Methods, value)
		return nil
	case FieldDefinition:
		td.Fields = append(td.Fields, value)
		return nil
	}
	return ErrNotApplicable
}

type IfDefinition struct {
	BaseDefinition
	Condition ConditionDefinition
	Body      Applicable
}

func (td IfDefinition) SetBody(body Applicable) {
	td.Body = body
}

func (td IfDefinition) Elem() interface{} {
	return td
}

func (td *IfDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case ConditionDefinition:
		td.Condition = value
		return nil
	case *ConditionDefinition:
		td.Condition = *value
		return nil
	}
	return ErrNotApplicable
}

const (
	Equal       Operator = iota + 1 // =
	NotEquality                     // !=
	Equality                        // ==
	Increment                       // ++
	Decrement                       // --
	Multiplication
	Subtraction
	Division
	Addition
	SelfMultiplication // *=
	SelfSubtraction // -=
	SelfDivision // -=
	SelfAddition // +=
	Modulo // %
	LessThan
	GreaterThan
	LessThanEqualTo
	GreaterThanEqualTo
	ConditionalAnd // &&
	ConditionalOR // ||
	BinaryAnd // &
	BinaryOR // |
	BitwiseNot // !
	BitwiseAnd // &
	BitwiseOR // !
	BitwiseXOR // ^
	LeftShift // <<
	RightShift // >>
)

type Operator int

type OperatorDefinition struct {
	BaseDefinition
	Operator Operator
}

func (td OperatorDefinition) Elem() interface{} {
	return td
}

func (td *OperatorDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case Operator:
		td.Operator = value
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type LoopDefinition struct {
	BaseDefinition
	Condition ConditionDefinition
	Body      Applicable
}

func (td LoopDefinition) SetBody(body Applicable) {
	td.Body = body
}

func (td LoopDefinition) Elem() interface{} {
	return td
}

func (td *LoopDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case ConditionDefinition:
		td.Condition = value
		return nil
	case *ConditionDefinition:
		td.Condition = *value
		return nil
	}
	return ErrNotApplicable
}

type ConditionDefinition struct {
	BaseDefinition
	Left     Applicable
	Right    Applicable
	Operator OperatorDefinition
}

func (td ConditionDefinition) Elem() interface{} {
	return td
}

func (td *ConditionDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type ForDefinition struct {
	BaseDefinition
	Left   Applicable
	Middle Applicable
	End    Applicable
	Body   Applicable
}

func (td ForDefinition) SetBody(body Applicable) {
	td.Body = body
}

func (td ForDefinition) Elem() interface{} {
	return td
}

func (td *ForDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type CaseDefinition struct {
	BaseDefinition
	Condition ConditionDefinition
	Body      Applicable
}

func (td CaseDefinition) SetBody(body Applicable) {
	td.Body = body
}

func (td CaseDefinition) Elem() interface{} {
	return td
}

func (td *CaseDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

type SwitchDefinition struct {
	BaseDefinition
	Cases     []CaseDefinition
	Condition ConditionDefinition
}

func (td SwitchDefinition) Elem() interface{} {
	return td
}

func (td *SwitchDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case ConditionDefinition:
		td.Condition = value
	case CaseDefinition:
		td.Cases = append(td.Cases, value)
	case []CaseDefinition:
		td.Cases = append(td.Cases, value...)
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	}
	return ErrNotApplicable
}

const (
	BothDirectional Direction = iota
	IncomingDirectional
	OutgoingDirectional
)

type Direction int

type ChannelDefinition struct {
	BaseDefinition
	Direction Direction
	Type      Applicable
}

func (td ChannelDefinition) Elem() interface{} {
	return td
}

func (td *ChannelDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case Direction:
		td.Direction = value
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case Applicable:
		td.Type = value
		return nil
	}
	return ErrNotApplicable
}

type FutureDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td FutureDefinition) Elem() interface{} {
	return td
}

func (td *FutureDefinition) Apply(item interface{}) error {
	if futureType, ok := item.(Applicable); ok {
		td.Type = futureType
	}
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case Applicable:
		td.Type = value
		return nil
	}
	return ErrNotApplicable
}

type StreamDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td StreamDefinition) Elem() interface{} {
	return td
}

func (td *StreamDefinition) Apply(item interface{}) error {
	switch value := item.(type) {
	case *BaseDefinition:
		td.BaseDefinition = *value
		return nil
	case BaseDefinition:
		td.BaseDefinition = value
		return nil
	case Applicable:
		td.Type = value
		return nil
	}
	return ErrNotApplicable
}
