package quiz

import (
	"github.com/gin-gonic/gin"
	"quizz-light/pkg/quiz/api"
	"quizz-light/pkg/quiz/repository"
)

type Controller struct {
	engine     *gin.Engine
	repository *repository.Repository
}

func NewController(basePathh string) *Controller {
	var repository repository.Repository = &repository.YamlFileRepository{
		BasePath: basePathh,
	}
	return &Controller{
		engine:     api.New(&repository),
		repository: &repository,
	}
}

func (c Controller) Run() {
	c.engine.Run()
}
