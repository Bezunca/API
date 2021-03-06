package parsers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

	return rawId, true
}

func ParseString(data map[string]interface{}, fieldName string) (string, bool) {

	fieldRaw, ok := data[fieldName]
	if !ok {
		return "", false
	}

	fieldString, ok := fieldRaw.(string)
	if !ok {
		return "", false
	}

	return fieldString, true
}

func ParseBool(data map[string]interface{}, fieldName string) (bool, bool) {

	fieldRaw, ok := data[fieldName]
	if !ok {
		return false, false
	}

	fieldBool, ok := fieldRaw.(bool)
	if !ok {
		return false, false
	}

	return fieldBool, true
}

func ParseInt32(data map[string]interface{}, fieldName string) (int32, bool) {

	fieldRaw, ok := data[fieldName]
	if !ok {
		return 0, false
	}

	fieldInt, ok := fieldRaw.(int32)
	if !ok {
		return 0, false
	}

	return fieldInt, true
}

func ParseTime(data map[string]interface{}, fieldName string) (primitive.DateTime, bool) {

	fieldRaw, ok := data[fieldName]
	if !ok {
		return primitive.NewDateTimeFromTime(time.Now()), false
	}

	fieldTime, ok := fieldRaw.(primitive.DateTime)
	if !ok {
		return primitive.NewDateTimeFromTime(time.Now()), false
	}

	return fieldTime, true
}
