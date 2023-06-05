package model

type Quiz interface {
	GetId() string
	GetTitle() string
	GetQuestions() []Question
}

type Question interface {
	GetText() string
	GetAnswers() []Answer
}

type Answer interface {
	GetText() string
	GetCorrect() bool
}
