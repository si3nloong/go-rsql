package rsql

import (
	"reflect"

	"github.com/timtadh/lexmachine"
)

const (
	defaultLimit = uint(50)
	// defaultMaxLimit = uint(100)
	maxUint = ^uint(0)
)

// FormatFunc :
type FormatFunc func(string) string

type Parser interface {
	ParseQuery(string) (*Params, error)
}

// RSQL :
type RSQL struct {
	SelectTag    string
	FilterTag    string
	SortTag      string
	LimitTag     string
	PageTag      string
	codec        *Struct
	lexer        *lexmachine.Lexer
	FormatColumn FormatFunc
	DefaultLimit uint
	MaxLimit     uint
}

// New :
func New(src interface{}) (*RSQL, error) {
	v := reflect.ValueOf(src)
	codec := getCodec(v.Type())
	lexer := lexmachine.NewLexer()
	dl := newDefaultTokenLexer()
	dl.addActions(lexer)

	p := new(RSQL)
	p.lexer = lexer
	p.FilterTag = "filter"
	p.SortTag = "sort"
	p.LimitTag = "limit"
	p.PageTag = "page"
	p.DefaultLimit = defaultLimit
	p.codec = codec
	return p, nil
}

// MustNew :
func MustNew(src interface{}) *RSQL {
	p, err := New(src)
	if err != nil {
		panic(err)
	}
	return p
}

// ParseQuery :
func (p *RSQL) ParseQuery(query string) (*Params, error) {
	return p.ParseQueryBytes([]byte(query))
}

// ParseQueryBytes :
func (p *RSQL) ParseQueryBytes(query []byte) (*Params, error) {
	values := make(map[string]string)
	if err := parseRawQuery(values, string(query)); err != nil {
		return nil, err
	}

	var (
		params = new(Params)
		// errs   = make(Errors, 0)
	)

	if err := p.parseFilter(values, params); err != nil {
		return nil, err
	}
	if err := p.parseSort(values, params); err != nil {
		return nil, err
	}
	if err := p.parseLimit(values, params); err != nil {
		return nil, err
	}
	if err := p.parseOffset(values, params); err != nil {
		return nil, err
	}

	return params, nil
}
