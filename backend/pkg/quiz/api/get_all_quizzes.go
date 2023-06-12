package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuizOverviewResponse struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (h handler) getAllQuizzes(ctx *gin.Context) {
	quizzes := h.Repository().GetAll()

	var quizOverviewResponse = make([]QuizOverviewResponse, len(quizzes))

	for index, quiz := range quizzes {
		quizOverviewResponse[index] = QuizOverviewResponse{
			Name:  quiz.GetId(),
			Title: quiz.GetTitle(),
		}
	}

	ctx.JSONP(http.StatusOK, &quizOverviewResponse)

}
