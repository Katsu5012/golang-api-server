package repository

import (
	model "golang-api-server/domain/models"
)

type QuestionRepository interface{
	Create(user * model.Question)(*model.Question,error)
	FindOne(id int)(*model.Question,error)
	Update(id int)(*model.Question,error)
	Delete(id int) error
}