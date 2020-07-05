package wallet

import (
	"bezuncapi/internal/database"
	"bezuncapi/internal/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateCEI(ctx echo.Context, email string, ceiCredentials models.CEI) bool {

	filter := bson.M{"auth_credentials.email": email}
	update := bson.D{
		{"$set", bson.D{
			{"wallets_credentials.cei", ceiCredentials},
		},
		},
	}

	updated := database.UpdateDocuments(ctx, database.UserDatabase, database.UsersCollection, filter, update)
	return updated
}
