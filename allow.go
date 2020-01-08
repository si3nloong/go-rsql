package rsql

import (
	"reflect"
	"time"
)

var (
	allowNums = []string{
		"eq", "ne",
		"gt", "gte", "lt", "lte",
		"in", "notIn",
		"between",
	}

	allows = map[interface{}][]string{
		reflect.String:              {"eq", "ne", "gt", "gte", "lt", "lte", "in", "notIn", "like", "notLike"},
		reflect.Bool:                {"eq", "ne"},
		reflect.Int:                 allowNums,
		reflect.Int8:                allowNums,
		reflect.Int16:               allowNums,
		reflect.Int32:               allowNums,
		reflect.Int64:               allowNums,
		reflect.Uint:                allowNums,
		reflect.Uint8:               allowNums,
		reflect.Uint16:              allowNums,
		reflect.Uint32:              allowNums,
		reflect.Uint64:              allowNums,
		reflect.Float32:             allowNums,
		reflect.Float64:             allowNums,
		reflect.TypeOf(time.Time{}): {},
	}
)
