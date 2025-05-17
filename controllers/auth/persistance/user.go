package persistance

import (
	"context"

	"github.com/FearLessSaad/SNFOK/db"
	auth_model "github.com/FearLessSaad/SNFOK/db/models/auth"
)

func GetUserByEmailAddress(email string) (bool, auth_model.Users) {

	conn := db.GetDB()
	ctx := context.Background()

	user := new(auth_model.Users)

	_ = conn.NewSelect().Model(user).Where("email = ?", email).Limit(1).Scan(ctx)

	if user.ID == "" {
		return false, auth_model.Users{}
	}

	return true, *user
}
