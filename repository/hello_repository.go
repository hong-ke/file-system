package repository

import "filesystem/entity"

type HelloRepository interface {
	Get() *entity.User
	Set(user *entity.User) error
}
