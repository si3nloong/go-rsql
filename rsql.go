package rsql

import (
	"log"
	"reflect"

	"github.com/timtadh/lexmachine"
)

const (
	defaultLimit    = uint(20)
	defaultMaxLimit = uint(100)
	maxUint         = ^uint(0)
)

// FormatFunc :
type FormatFunc func(string) string

// Parser :
type Parser struct {
	SelectTag    string
	FilterTag    string
	SortTag      string
	LimitTag     string
	codec        *Struct
	lexer        *lexmachine.Lexer
	FormatColumn FormatFunc
	DefaultLimit uint
	MaxLimit     uint
	// registry     *codec.Registry
	// mapper       *reflext.Struct
}

// New :
func New(src interface{}) (*Parser, error) {
	v := reflect.ValueOf(src)
	log.Println(v, v.Type())
	codec := getCodec(v.Type())
	log.Println(codec)
	lexer := lexmachine.NewLexer()
	dl := newDefaultTokenLexer()
	dl.addActions(lexer)

	p := new(Parser)
	p.lexer = lexer
	p.FilterTag = "filter"
	p.codec = codec
	return p, nil
}

// ParseQuery :
func (p *Parser) ParseQuery(query string) (*Params, error) {
	return p.ParseQueryBytes([]byte(query))
}

// ParseQueryBytes :
func (p *Parser) ParseQueryBytes(query []byte) (*Params, error) {
	values := make(map[string]string)
	if err := parseRawQuery(values, string(query)); err != nil {
		return nil, err
	}

	var (
		params = new(Params)
		// errs   = make(Errors, 0)
	)

	// errs = append(errs, p.parseSelect(values, params)...)
	// errs = append(errs, p.parseSort(values, params)...)
	// errs = append(errs, p.parseLimit(values, params)...)
	if err := p.parseFilter(values, params); err != nil {
		return nil, err
	}

	// if len(errs) > 0 {
	// 	return nil, errs
	// }
	return params, nil
}
