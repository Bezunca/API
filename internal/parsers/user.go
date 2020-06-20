package parsers

import "bezuncapi/internal/models"

func ParseLoginCredentials(data map[string]interface{}) (models.LoginCredentials, bool) {

	_loginCredentials, ok := data["login_credentials"]
	if !ok {
		return models.LoginCredentials{}, false
	}

	loginCredentials, ok := _loginCredentials.(map[string]interface{})
	if !ok {
		return models.LoginCredentials{}, false
	}

	return models.LoginCredentials{
		Email: loginCredentials["email"].(string),
	}, true
}

func ParseUser(data map[string]interface{}) (models.User, bool) {

	id, ok := ParseID(data)
	if !ok {
		return models.User{}, false
	}

	loginCredentials, ok := ParseLoginCredentials(data)
	if !ok {
		return models.User{}, false
	}

	return models.User{
		ID:               id,
		LoginCredentials: loginCredentials,
	}, true
}
