package custom_scalars

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"time"
)


func MarshalDateTime(b time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(b.Format("2006-01-02 15:04:05")))
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		dateTime, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil{
			return time.Time{}, err
		}
		return dateTime, nil
	default:
		return time.Time{}, fmt.Errorf("%T is not a DateTime", v)
	}
}