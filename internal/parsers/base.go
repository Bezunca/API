package parsers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseID(data map[string]interface{}) (primitive.ObjectID, bool) {
	_rawId, ok := data["_id"]
	if !ok {
		return primitive.ObjectID{}, false
	}

	rawId, ok := _rawId.(primitive.ObjectID)
	if !ok {
		return primitive.ObjectID{}, false
	}

	/*id := make([]uint8, len(rawId))
	for i := 0; i < len(id); i++ {
		id[i] = rawId[i]
	}*/

	return rawId, true
}
