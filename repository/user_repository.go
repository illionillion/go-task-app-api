package repository

import "go-rest-api/model"

// ユーザーリポジトリのインターフェース
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}