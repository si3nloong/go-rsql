package rsql

// Params :
type Params struct {
	Selects []string
	Filters Filters
	Sorts   []Sort
	Limit   uint
	Offset  uint
	Cursor  string
}

type Filters []Filter

func (fs Filters) Lookup(key string) (val interface{}, ok bool) {
	for _, f := range fs {
		if key == f.Name {
			val = f.Value
			ok = true
			return
		}
	}
	ok = false
	return
}
