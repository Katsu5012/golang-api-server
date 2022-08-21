package repository

import (
	model "golang-api-server/domain/models"
)

type UserRepositoy interface{
	Create(user * model.User)(*model.User,error)
	FindOne(id int)(*model.User,error)
	Update(user * model.User)(*model.User,error)
	Delete(user *model.User) error
}