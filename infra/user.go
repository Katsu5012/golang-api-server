package infra

import (
	model "golang-api-server/domain/models"
	"golang-api-server/domain/repository"

	"github.com/jinzhu/gorm"
)

type UserRepository struct{
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepositoy{
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) Create(user *model.User)(*model.User,error){
     if err :=ur.Conn.Create(&user).Error ; err !=nil{
		 return nil,err
	 }

	 return user,nil
}

func (ur *UserRepository) FindOne(id int)(*model.User,error){
	user :=&model.User{ID:id}

	if err :=ur.Conn.First(&user).Error;err !=nil{
		return nil,err
	}

	return user,nil
}

func (ur * UserRepository) Update(user *model.User)(*model.User,error){
	if err :=ur.Conn.Model(&user).Update(&user).Error;err !=nil{
		return nil,err
	}

	return user,nil
}

func (ur *UserRepository) Delete(user *model.User)error{
	if err :=ur.Conn.Delete(&user).Error;err !=nil{
		return err
	}

	return nil
}