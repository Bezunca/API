package parsers

import (
	"github.com/Bezunca/API/internal/models"
)

func ParseWalletsCredentials(data map[string]interface{}) (models.WalletsCredentials, bool) {

	_walletsCredentials, ok := data["wallets_credentials"]
	if !ok {
		return models.WalletsCredentials{}, false
	}

	walletsCredentials, ok := _walletsCredentials.(map[string]interface{})
	if !ok {
		return models.WalletsCredentials{}, false
	}

	cei, ok := ParseCEI(walletsCredentials)
	if !ok {
		cei = nil
	}

	return models.WalletsCredentials{
		Cei: cei,
	}, true
}

func ParseCEI(data map[string]interface{}) (*models.CEI, bool) {

	_cei, ok := data["cei"]
	if !ok {
		return nil, false
	}

	cei, ok := _cei.(map[string]interface{})
	if !ok {
		return nil, false
	}

	user, ok := ParseString(cei, "user")
	if !ok {
		return nil, false
	}

	password, ok := ParseString(cei, "password")
	if !ok {
		return nil, false
	}

	status, ok := ParseSyncStatus(cei)
	if !ok {
		return nil, false
	}

	return &models.CEI{
		User:     user,
		Password: password,
		Status:   status,
	}, true
}

func ParseSyncStatus(data map[string]interface{}) (models.SyncStatus, bool) {

	_status, ok := data["status"]
	if !ok {
		return models.SyncStatus{}, false
	}

	status, ok := _status.(map[string]interface{})
	if !ok {
		return models.SyncStatus{}, false
	}

	statusType, ok := ParseInt32(status, "status_type")
	if !ok {
		return models.SyncStatus{}, false
	}

	statusMessage, ok := ParseString(status, "status_message")
	if !ok {
		return models.SyncStatus{}, false
	}

	statusDate, ok := ParseTime(status, "status_date")
	if !ok {
		return models.SyncStatus{}, false
	}

	return models.SyncStatus{
		StatusType: statusType,
		StatusMessage: statusMessage,
		StatusDate: statusDate,
	}, true
}
