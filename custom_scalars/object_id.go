package custom_scalars

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)


func MarshalObjectID(b primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(b.Hex()))
	})
}

func UnmarshalObjectID(v interface{}) (primitive.ObjectID, error) {
	switch v := v.(type) {
	case string:
		oID, err := primitive.ObjectIDFromHex(v)
		if err != nil{
			return primitive.ObjectID{}, err
		}
		return oID, nil
	default:
		return primitive.ObjectID{}, fmt.Errorf("%T is not an ObjectID", v)
	}
}