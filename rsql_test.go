package rsql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestRSQL :
func TestRSQL(t *testing.T) {
	var i struct {
		Name   string  `rsql:"n,filter,sort,allow=eq|gt|gte"`
		Status string  `rsql:"status,filter,sort"`
		PtrStr *string `rsql:"text,filter,sort"`
		No     int     `rsql:"no,column=No2,filter,sort"`
	}

	p := MustNew(i)

	{
		param, err := p.ParseQuery(`filter=status=eq="111";no=gt=1991;text==null&sort=status,-no&limit=100`)
		require.NoError(t, err)
		require.NotNil(t, param)
		require.True(t, len(param.Filters) > 0)
		require.True(t, len(param.Sorts) > 0)
		require.Equal(t, uint(100), param.Limit)
	}
}
