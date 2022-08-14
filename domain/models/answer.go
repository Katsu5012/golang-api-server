package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Answer struct{
	gorm.Model
	ID int `gorm:"primaryKey"`
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	QuestionID int
	Question Question `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewAnswer(content string)(*Question,error){
	if content == ""{
		return nil,errors.New("内容が入力されていません")
	}

	question :=&Question{
	   Content:content,
	} 

	return question,nil
}