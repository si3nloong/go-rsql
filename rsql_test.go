package rsql

import (
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
			No          int       `rsql:"no,filter,sort,column=No2"`
			Int         CustomInt `rsql:"int,filter"`
			SubmittedAt time.Time `rsql:"submittedAt,filter"`
			CreatedAt   time.Time `rsql:"createdAt,sort"`
		}

		p := MustNew(i)
		param, err := p.ParseQuery(`filter=int>10;status=eq="111";no=gt=1991;text==null&sort=status,-no&limit=100&page=2`)
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
		// log.Println("Filters :", param.Filters)

		param, err = p.ParseQuery(`filter=(submittedAt>='2019-12-22T16:00:00Z';submittedAt<='2019-12-31T15:59:59Z';status=="APPROVED")&limit=100`)
		require.NoError(t, err)
		require.NotNil(t, param)
		// log.Println("Filters :", param.Filters)
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

	{
		var i struct {
			Title            string    `rsql:"title,filter"`
			Audience         string    `rsql:"audience,filter"`
			Status           string    `rsql:"status,filter,sort,allow=eq|gt|gte"`
			ScheduleDateTime time.Time `rsql:"scheduleDateTime,filter,sort"`
		}
		p := MustNew(i)
		param, err := p.ParseQuery(`filter=(audience=="CUSTOMIZED";status=="PENDING";scheduleDateTime>='2020-01-14T16:00:00Z';scheduleDateTime<='2020-01-20T15:59:59Z';title=like="testing%25")`)
		require.NoError(t, err)
		require.NotNil(t, param)
	}

	{
		var i struct {
			Name             string    `rsql:"name,sort"`
			Status           string    `rsql:"status,filter,sort,allow=eq|gt|gte"`
			ScheduleDateTime time.Time `rsql:"scheduleDateTime,filter,sort"`
		}

		p := MustNew(i)
		query := `filter=&sort=name,-status&limit=10&page=2`
		param, err := p.ParseQuery(query)
		require.NoError(t, err)
		require.Equal(t, uint(10), param.Offset)

		/*
			actions.Find().
				From("table").
				Where(
					expr.Equal("key", "value"),
				).
				OrderBy(
					expr.Asc("F1"),
					expr.Asc("F2"),
					expr.Desc("Status"),
				).
				Limit(10).
				Offset(2 * 10)
		*/

	}
}
