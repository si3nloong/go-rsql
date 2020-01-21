package rsql

import (
	"fmt"
	"strconv"
)

func (p *RSQL) parseLimit(values map[string]string, params *Params) error {
	val, ok := values[p.LimitTag]
	params.Limit = p.DefaultLimit
	delete(values, p.LimitTag)
	if !ok || len(val) < 1 {
		return nil
	}

	u64, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	if u64 > uint64(maxUint) {
		return fmt.Errorf("overflow unsigned integer, %d", u64)
	}
	params.Limit = uint(u64)
	return nil
}

func (p *RSQL) parseOffset(values map[string]string, params *Params) error {
	val, ok := values[p.PageTag]
	delete(values, p.PageTag)
	if !ok || len(val) < 1 {
		return nil
	}

	u64, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}
	if u64 > uint64(maxUint) {
		return fmt.Errorf("overflow unsigned integer, %d", u64)
	}
	params.Offset = uint(u64) * params.Limit
	return nil
}
