package rsql

// Params :
type Params struct {
	Selects []string
	Filters []*Filter
	Sorts   []*Sort
	Limit   uint
	Offset  uint
	Cursor  string
}
