package usecase

import (
	model "golang-api-server/domain/models"
	"golang-api-server/domain/repository"
)

type UserUsecase interface{
	Create(name,mail string,isAdmin bool)(*model.User,error)
	FindOne(id int)(*model.User,error)
	Update(id int)(*model.User,error)
	Delete(id int)error
}

type UserUsercase struct{
	userRepository repository.UserRepositoy
}

func NewUserUsecase(ur repository.UserRepositoy)UserUsecase{
	return &ur(userRepository:ur)
}