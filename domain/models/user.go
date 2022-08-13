package model

import "errors"


type User struct{
	Id int 
	Name string
	MailAddress string	
	IsAdminUser bool
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




