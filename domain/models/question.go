package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Question struct{
	gorm.Model
	ID int 
	Title string
	Content string
    UserId uint
    CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewQuestion(title,content string,isAnswer bool)(*Question,error){
	if title == "" {
		return nil,errors.New("タイトルが入力されていません")
	}

	if content == ""{
		return nil,errors.New("内容が入力されていません")
	}

	question :=&Question{
       Title :title,
	   Content:content,
	}

	return question,nil
}