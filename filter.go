package rsql

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/timtadh/lexmachine"
)

// Filter :
type Filter struct {
	Name     string
	Operator Expr
	Value    interface{}
}

func (p *Parser) parseFilter(values map[string]string, params *Params) error {
	val, ok := values[p.FilterTag]
	if !ok || len(val) < 1 {
		return nil
	}

	scan, err := p.lexer.Scanner([]byte(val))
	if err != nil {
		return err
	}

	conds := make([]*Filter, 0)
loop:
	for {
		tkn1, err := nextToken(scan)
		if err != nil {
			if err == io.EOF {
				break loop
			}
			return err
		}

		f, ok := p.codec.Names[tkn1.Value]
		if !ok {
			return fmt.Errorf("invalid field to filter")
		}

		if _, ok := f.Tag.Lookup("filter"); !ok {
			return fmt.Errorf("invalid field to filter")
		}

		filters := allows[f.Type.Kind()]
		if v, ok := f.Tag.Lookup("allow"); ok {
			filters = strings.Split(v, "|")
		}

		log.Println("debug /...................")
		log.Println(filters, f.Type, tkn1)
		tkn2, err := nextToken(scan)
		if err != nil {
			return err
		}
		tkn3, err := nextToken(scan)
		if err != nil {
			return err
		}

		value, err := convertValue(f, strings.Trim(tkn3.Value, `"`))
		if err != nil {
			return err
		}

		conds = append(conds, &Filter{
			Name:     tkn1.Value,
			Operator: operators[tkn2.Value],
			Value:    value,
		})

		tkn, err := nextToken(scan)
		if err != nil {
			if err == io.EOF {
				break loop
			}
			return err
		}

		switch tkn.Value {
		case ";", ",":
		default:
			return errors.New("unexpected char")
		}
	}

	for _, c := range conds {
		log.Println("Each :", *c, reflect.TypeOf(c.Value))
	}
	return nil
}

func nextToken(scan *lexmachine.Scanner) (*Token, error) {
	it, err, eof := scan.Next()
	if eof {
		return nil, io.EOF
	}
	if err != nil {
		return nil, err
	}
	return it.(*Token), nil
}

func convertValue(sf *StructField, value string) (interface{}, error) {
	v := reflect.New(sf.Type).Elem()

	switch sf.Type.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Bool:
		x, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		v.SetBool(x)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, err
		}
		if v.OverflowInt(x) {
			return nil, errors.New("int overflow")
		}
		v.SetInt(x)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return nil, err
		}
		if v.OverflowUint(x) {
			return nil, errors.New("unsigned int overflow")
		}
		v.SetUint(x)
	case reflect.Float32, reflect.Float64:
		x, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		if v.OverflowFloat(x) {
			return nil, errors.New("float overflow")
		}
		v.SetFloat(x)
	case reflect.Array:
	case reflect.Slice:
	case reflect.Ptr:
	}

	return v.Interface(), nil
}
