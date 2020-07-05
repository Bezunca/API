package parsers

import (
	"bezuncapi/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseAuthCredentials(data map[string]interface{}) (models.AuthCredentials, bool) {

	_authCredentials, ok := data["auth_credentials"]
	if !ok {
		return models.AuthCredentials{}, false
	}

	authCredentials, ok := _authCredentials.(map[string]interface{})
	if !ok {
		return models.AuthCredentials{}, false
	}

	email, ok := ParseString(authCredentials, "email")
	if !ok {
		return models.AuthCredentials{}, false
	}

	password, ok := ParseString(authCredentials, "password")
	if !ok {
		return models.AuthCredentials{}, false
	}

	activated, ok := ParseBool(authCredentials, "activated")
	if !ok {
		return models.AuthCredentials{}, false
	}

	return models.AuthCredentials{
		Email:     email,
		Password:  password,
		Activated: activated,
	}, true
}

func ParseUsers(cursor *mongo.Cursor) (interface{}, bool) {

	var users []models.User

	for cursor.Next(context.TODO()) {

		var data map[string]interface{}
		err := cursor.Decode(&data)
		if err != nil {
			return nil, false
		}

		id, ok := ParseID(data)
		if !ok {
			return nil, false
		}

		authCredentials, ok := ParseAuthCredentials(data)
		if !ok {
			return nil, false
		}

		name, ok := ParseString(data, "name")
		if !ok {
			return nil, false
		}

		user := models.User{
			ID:              id,
			Name:            name,
			AuthCredentials: authCredentials,
		}

		users = append(users, user)
	}

	return users, true
}
