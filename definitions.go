package rewrite

type Meta struct {
	Version string `json:"version"`
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

func (bd *BaseDefinition) SetDescription(desc string) {
	bd.Description = desc
}

func (bd *BaseDefinition) SetName(name string) {
	bd.Name = name
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

// TypeDefinition defines the base definition for types.
type TypeDefinition struct {
	BaseDefinition
	Type   BaseType
	Memory MemoryLayout
}

func (td TypeDefinition) Elem() interface{} {
	return td
}

type CallDefinition struct {
	BaseDefinition
	Target    Applicable
	Arguments []VariableDefinition
}

func (td CallDefinition) Elem() interface{} {
	return td
}

type CommentDefinition struct {
	BaseDefinition
	Contents []string
}

func (td CommentDefinition) Elem() interface{} {
	return td
}

type AnnotationDefinition struct {
	BaseDefinition
	Content string
}

func (td AnnotationDefinition) Elem() interface{} {
	return td
}

type ResultDefinition struct {
	BaseDefinition
	Type  Applicable
	Value Applicable
}

func (td ResultDefinition) Elem() interface{} {
	return td
}

type Value struct {
	BaseDefinition
	Value interface{}
}

func (td Value) Elem() interface{} {
	return td
}

type AssignmentDefinition struct {
	BaseDefinition
	Value Applicable
	Short bool // should use := instead of =
}

func (td AssignmentDefinition) Elem() interface{} {
	return td
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

type ReturnDefinition struct {
	BaseDefinition
	Targets []Applicable
}

func (td ReturnDefinition) Elem() interface{} {
	return td
}

type FieldDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td FieldDefinition) Elem() interface{} {
	return td
}

type LiteralDefinition struct {
	BaseDefinition
	Literal string
}

func (td LiteralDefinition) Elem() interface{} {
	return td
}

// DataType represents a defined type where the type is a previously
// created data type.
type DataTypeDefinition struct {
	BaseDefinition
	Type string
}

func (td DataTypeDefinition) Elem() interface{} {
	return td
}

type DataDefinition struct {
	BaseDefinition
	Fields  []FieldDefinition
	Methods []Function
}

func (td DataDefinition) Elem() interface{} {
	return td
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

const (
	Equal       Operator = iota + 1 // =
	ShortEqual                      // :=
	NotEquality                     // !=
	Equality                        // ==
	Increment                       // ++
	Decrement                       // --
	Multiplication
	Subtraction
	Division
	Addition
	SelfMultiplication // *=
	SelfSubtraction    // -=
	SelfDivision       // -=
	SelfAddition       // +=
	Modulo             // %
	LessThan
	GreaterThan
	LessThanEqualTo
	GreaterThanEqualTo
	ConditionalAnd  // &&
	ConditionalOR   // ||
	BinaryAnd       // &
	BinaryOR        // |
	BitwiseNot      // ~
	BitwiseAnd      // &
	BitwiseOR       // |
	BitwiseXOR      // ^
	LeftShift       // <<
	RightShift      // >>
	PointerAnd      // &
	PointerAsterick // *
)

type Operator int

type OperatorDefinition struct {
	BaseDefinition
	Operator Operator
}

func (td OperatorDefinition) Elem() interface{} {
	return td
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

type ForDefinition struct {
	BaseDefinition
	Start  Applicable
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

type SwitchDefinition struct {
	BaseDefinition
	Cases     []CaseDefinition
	Condition ConditionDefinition
}

func (td SwitchDefinition) Elem() interface{} {
	return td
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

type FutureDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td FutureDefinition) Elem() interface{} {
	return td
}

type Function struct {
	BaseDefinition
	Pointer bool
	Returns []ResultDefinition
	Accepts []VariableDefinition
	Body    []Applicable
}

func (td Function) Elem() interface{} {
	return td
}

type AsyncFunction struct {
	Function
}

type StreamDefinition struct {
	BaseDefinition
	Type Applicable
}

func (td StreamDefinition) Elem() interface{} {
	return td
}
