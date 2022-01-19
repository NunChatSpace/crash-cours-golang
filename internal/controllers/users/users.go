package users

import (
	"context"

	"github.com/NunChatSpace/crash-cours-golang/internal/models"
)

func getTempData() []models.User {
	return []models.User{
		{
			FirstName: "nun",
			LastName:  "test",
		},
		{
			FirstName: "nun2",
			LastName:  "test",
		},
	}
}

func GetUsers(c context.Context) []models.User {
	return getTempData()
}
