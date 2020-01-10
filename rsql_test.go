package rsql

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type CustomInt int

// TestRSQL :
func TestRSQL(t *testing.T) {
	{
		var i struct {
			Name        string    `rsql:"n,filter,sort,allow=eq|gt|gte"`
			Status      string    `rsql:"status,filter,sort"`
			PtrStr      *string   `rsql:"text,filter,sort"`
			No          int       `rsql:"no,column=No2,filter,sort"`
			Int         CustomInt `rsql:"int,filter"`
			SubmittedAt time.Time `rsql:"submittedAt,filter"`
			CreatedAt   time.Time `rsql:"createdAt,sort"`
		}

		p := MustNew(i)
		param, err := p.ParseQuery(`filter=int>10;status=eq="111";no=gt=1991;text==null&sort=status,-no&limit=100`)
		require.NoError(t, err)
		require.NotNil(t, param)
		require.True(t, len(param.Filters) > 0)
		require.True(t, len(param.Sorts) > 0)
		require.Equal(t, uint(100), param.Limit)

		param, err = p.ParseQuery(`filter=(int>10;status=eq="111";no=gt=1991;text==null)&sort=status,-no&limit=100`)
		require.NoError(t, err)
		require.NotNil(t, param)

		param, err = p.ParseQuery(`filter=(status=="APPROVED";submittedAt>="2019-12-22T16:00:00Z";submittedAt<="2019-12-31T15:59:59Z")&sort=-createdAt&limit=10`)
		require.NoError(t, err)
		require.NotNil(t, param)
		log.Println("Filters :", param.Filters)

		param, err = p.ParseQuery(`filter=(submittedAt>='2019-12-22T16:00:00Z';submittedAt<='2019-12-31T15:59:59Z';status=='APPROVED')&limit=100`)
		require.NoError(t, err)
		require.NotNil(t, param)
		log.Println("Filters :", param.Filters)
	}

	{
		var i struct {
			Flag        bool      `rsql:"flag,filter"`
			Status      string    `rsql:"status,filter,sort,allow=eq|gt|gte"`
			SubmittedAt time.Time `rsql:"submittedAt,filter,sort"`
		}
		p := MustNew(i)
		param, err := p.ParseQuery(`filter=status=eq="approved";flag=ne=false;submittedAt>="2019-12-22T16:00:00Z";submittedAt<="2019-12-31T15:59:59Z"&sort=status&limit=10`)
		require.NoError(t, err)
		require.NotNil(t, param)
	}
}
