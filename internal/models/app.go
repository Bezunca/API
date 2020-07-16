package models

type AppInfo struct {
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Cei   *SyncStatus `json:"cei"`
}
