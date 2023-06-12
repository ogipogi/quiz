package repository

import "quizz-light/pkg/quiz/model"

type Repository interface {
	GetAll() []model.Quiz
	Get(id string) model.Quiz
}
