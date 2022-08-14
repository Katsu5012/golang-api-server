package repository

import (
	model "golang-api-server/domain/models"
)

type UserRepositoy interface{
	Create(user * model.User)(*model.User,error)
	FindOne(id int)(*model.User,error)
	Update(id int)(*model.User,error)
	Delete(id int) error
}