package service

import (
	"context"
	"golang-database-user/model"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.MstUser, roleId string) model.MstUser
	UpdateUser(ctx context.Context, user model.MstUser, userId string) model.MstUser
	DeleteUser(ctx context.Context, userId string) (model.MstUser, error)
	ReadUsers(ctx context.Context) ([]model.MstUser, error)
}
