package rsql

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Direction int

const (
	Asc Direction = iota
	Desc
)

// Sort :
type Sort struct {
	Field     string
	Direction Direction
}

func (p *RSQL) parseSort(values map[string]string, params *Params) error {
	val, ok := values[p.SortTag]
	delete(values, p.SortTag)
	if !ok || len(val) < 1 {
		return nil
	}

	paths := strings.Split(val, ",")
	for _, v := range paths {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			return errors.New("rsql: invalid sort")
		}

		v, err := url.QueryUnescape(v)
		if err != nil {
			return err
		}

		dir := Asc
		desc := v[0] == '-'
		if desc {
			v = v[1:]
			dir = Desc
		}

		f, ok := p.codec.Names[v]
		if !ok {
			return fmt.Errorf("rsql: invalid field %q to sort", v)
		}

		if _, ok := f.Tag.Lookup("sort"); !ok {
			return fmt.Errorf("rsql: field %q is not allow to sort", v)
		}

		name := f.Name
		if v, ok := f.Tag.Lookup("column"); ok {
			name = v
		}

		params.Sorts = append(params.Sorts, Sort{
			Field:     name,
			Direction: dir,
		})
	}
	return nil
}
