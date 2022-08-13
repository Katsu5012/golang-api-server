package model

import "errors"

type Question struct{
	Id int `validate:"required"`
	Title string
	Content string
	IsAnswer bool
	RespondentId int
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