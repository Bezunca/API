package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WalletsCredentials struct {
	Cei *CEI `bson:"cei" json:"cei"`
}

type CEI struct {
	User     string     `bson:"user" json:"user"`
	Password string     `bson:"password" json:"password"`
	Status   SyncStatus `bson:"status" json:"status"`
}

type SyncStatus struct {
	StatusType    int32              `bson:"status_type" json:"status_type"`
	StatusMessage string             `bson:"status_message" json:"status_message"`
	StatusDate    primitive.DateTime `bson:"status_date" json:"status_date"`
}

const (
	StatusPending = 1
	StatusUnauthorized = 2
	StatusSyncError = 3
	StatusSyncOk = 4
)