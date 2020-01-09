package rsql

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRSQL :
func TestRSQL(t *testing.T) {
	var i struct {
		Name   string  `rsql:"n,filter,sort,allow=eq|gt|gte"`
		Status string  `rsql:"status,filter"`
		PtrStr *string `rsql:"text,filter"`
		No     int     `rsql:"no,column=No2,filter"`
	}

	p, err := New(i)
	log.Println(p, err)
	params, err := p.ParseQuery(`filter=status=eq="111";no=gt=1991;text==null`)
	require.NoError(t, err)
	log.Println(params, err)
}
