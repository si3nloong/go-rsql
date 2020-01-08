package rsql

import (
	"log"
	"testing"
)

// TestRSQL :
func TestRSQL(t *testing.T) {
	var i struct {
		Name   string `rsql:"n,filter,sort,allows=eq|gt|gte"`
		Status string `rsql:"status,filter"`
		No     int    `rsql:"no,column=No2,filter"`
	}

	p, err := New(i)
	log.Println(p, err)
	p.ParseQuery(`filter=status=eq="111";no=gt=1991`)
}
