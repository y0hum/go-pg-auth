package store

import "github.com/ZeroDayDrake/go-pg-auth/src/http/models"

type UserRepository interface {
	Create(*models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
}