package rsql

// Expr :
type Expr int

// types :
const (
	Equal Expr = iota
	NotEqual
	LesserThan
	LesserOrEqual
	GreaterThan
	GreaterOrEqual
	Like
	NotLike
	In
	NotIn
)

func (e Expr) String() string {
	switch e {
	case Equal:
		return "eq"
	case NotEqual:
		return "ne"
	case GreaterThan:
		return "gt"
	case GreaterOrEqual:
		return "gte"
	case LesserThan:
		return "lt"
	case LesserOrEqual:
		return "lte"
	case In:
		return "in"
	case NotIn:
		return "notIn"
	default:
		return "unknown"
	}
}

var operators = map[string]Expr{
	"==":      Equal,
	"=eq=":    Equal,
	"!=":      NotEqual,
	"=ne=":    NotEqual,
	">":       GreaterThan,
	"=gt=":    GreaterThan,
	">=":      GreaterOrEqual,
	"=gte=":   GreaterOrEqual,
	"<":       LesserThan,
	"=lt=":    LesserThan,
	"<=":      LesserOrEqual,
	"=lte=":   LesserOrEqual,
	"=like=":  Like,
	"=nlike=": NotLike,
	"=in=":    In,
	"=nin=":   NotIn,
}
