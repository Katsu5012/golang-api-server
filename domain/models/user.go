package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)


type User struct{
	gorm.Model
	ID int 
	Name string
	// TODO: ValueObjectを作成する
	MailAddress string	
	// NOTE: Rollにするか検討する
	IsAdminUser bool
	Questions []Question
	Answers []Answer
    CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Userのconstructor
func NewUser(name,mail string,isAdmin bool) (*User,error){
if name == ""{
	return nil,errors.New("名前が入力されていません")
}

user := &User{
	Name: name,
	MailAddress: mail,
	IsAdminUser: isAdmin,
}

return user,nil
}




