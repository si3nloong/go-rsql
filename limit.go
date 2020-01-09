package rsql

import (
	"fmt"
	"strconv"
)

func (p *Parser) parseLimit(values map[string]string, params *Params) error {
	val, ok := values[p.LimitTag]
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
	// if params.Limit > p.MaxLimit {

	// }
	return nil
}
