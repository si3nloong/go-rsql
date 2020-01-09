package rsql

// Params :
type Params struct {
	Selects []string
	Filters []*Filter
	Sorts   []*Sort
	Limit   uint
	Cursor  string
}
