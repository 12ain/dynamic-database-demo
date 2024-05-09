package repository

import "dynamic-database-demo/model"

type UserRepository interface {
	CreateUser(user model.User) error
	GetUser(userID string) (model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(userID string) error
}
