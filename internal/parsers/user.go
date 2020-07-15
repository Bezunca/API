package parsers

import (
	"context"
	"github.com/Bezunca/API/internal/models"
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

func ParseCEI(data map[string]interface{}) (models.CEI, bool) {

	_cei, ok := data["cei"]
	if !ok {
		return models.CEI{}, false
	}

	cei, ok := _cei.(map[string]interface{})
	if !ok {
		return models.CEI{}, false
	}

	user, ok := ParseString(cei, "user")
	if !ok {
		return models.CEI{}, false
	}

	password, ok := ParseString(cei, "password")
	if !ok {
		return models.CEI{}, false
	}

	return models.CEI{
		User: user,
		Password: password,
	}, true
}

func ParseWalletsCredentials(data map[string]interface{}) (models.WalletCredentials, bool) {

	_walletsCredentials, ok := data["wallets_credentials"]
	if !ok {
		return models.WalletCredentials{}, false
	}

	walletsCredentials, ok := _walletsCredentials.(map[string]interface{})
	if !ok {
		return models.WalletCredentials{}, false
	}

	cei, ok := ParseCEI(walletsCredentials)
	if !ok {
		cei = models.CEI{}
	}

	return models.WalletCredentials{
		Cei: cei,
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

		name, ok := ParseString(data, "name")
		if !ok {
			return nil, false
		}

		authCredentials, ok := ParseAuthCredentials(data)
		if !ok {
			return nil, false
		}

		walletsCredentials, ok := ParseWalletsCredentials(data)
		if !ok {
			walletsCredentials = models.WalletCredentials{}
		}

		user := models.User{
			ID:              id,
			Name:            name,
			AuthCredentials: authCredentials,
			WalletsCredentials: walletsCredentials,
		}

		users = append(users, user)
	}

	return users, true
}
