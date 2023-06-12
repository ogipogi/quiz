package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"quizz-light/pkg/quiz/repository"
)

type handler struct {
	repository *repository.Repository
}

func (h handler) Repository() repository.Repository {
	return *h.repository
}

func New(repository *repository.Repository) *gin.Engine {
	engine := gin.Default()

	h := &handler{repository: repository}

	configureCors(engine)

	routes := engine.Group("/quizzes")
	routes.GET("/all", h.getAllQuizzes)

	return engine
}

func configureCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET"}
	engine.Use(cors.New(config))
}
