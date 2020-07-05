package models

type WalletCredentials struct {
	Cei CEI `bson:"cei" json:"cei"`
}

type CEI struct {
	User       string `bson:"user" json:"user"`
	Password  string `bson:"password" json:"password"`
}
