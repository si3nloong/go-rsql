package rsql

// Params :
type Params struct {
	Selects []interface{}
	Filters []*Filter
	Sorts   []*Sort
	Limit   uint
}
