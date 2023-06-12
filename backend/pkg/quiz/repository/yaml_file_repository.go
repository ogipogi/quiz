package repository

import (
	"gopkg.in/yaml.v3"
	"os"
	"quizz-light/pkg/quiz/model"
	"quizz-light/pkg/utils/log"
	"strings"
)

type YamlFileRepository struct {
	BasePath string
}

type QuizYaml struct {
	Id        string
	Title     string         `yaml:"title"`
	Questions []QuestionYaml `yaml:"questions"`
}

type QuestionYaml struct {
	Text    string       `yaml:"question"`
	Answers []AnswerYaml `yaml:"answers"`
}

type AnswerYaml struct {
	Text    string `yaml:"text"`
	Correct bool   `yaml:"correct,omitempty"`
}

func (q *QuizYaml) GetId() string {
	return q.Id
}

func (q *QuizYaml) GetTitle() string {
	return q.Title
}

func (q *QuizYaml) GetQuestions() []model.Question {
	questions := make([]model.Question, len(q.Questions))
	for index, question := range q.Questions {
		q := new(QuestionYaml)
		*q = question
		questions[index] = q
	}
	return questions
}

func (q *QuestionYaml) GetText() string {
	return q.Text
}

func (q *QuestionYaml) GetAnswers() []model.Answer {
	answers := make([]model.Answer, len(q.Answers))
	for index, answer := range q.Answers {
		a := new(AnswerYaml)
		*a = answer
		answers[index] = a
	}
	return answers
}

func (a *AnswerYaml) GetText() string {
	return a.Text
}

func (a *AnswerYaml) GetCorrect() bool {
	return a.Correct
}

func (repository YamlFileRepository) GetAll() []model.Quiz {
	dir, _ := os.ReadDir(repository.BasePath)
	quizzes := make([]model.Quiz, len(dir))
	for index, entry := range dir {
		quiz := repository.Get(strings.TrimSuffix(entry.Name(), ".yaml"))
		quizzes[index] = quiz
	}
	return quizzes
}

func (repository YamlFileRepository) Get(id string) model.Quiz {
	file := readFile(repository.BasePath, id)
	quiz := &QuizYaml{}
	err := yaml.Unmarshal(file, quiz)
	if err != nil {
		log.Error("could not unmarshal:", id)
	}
	quiz.Id = id
	return quiz
}

func readFile(basePath, fileName string) []byte {
	absoluteFilePath := basePath + "/" + fileName + ".yaml"
	file, err := os.ReadFile(absoluteFilePath)
	if err != nil {
		log.Error("could not read file:", absoluteFilePath)
		return []byte{}
	}
	return file
}
