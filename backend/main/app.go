package main

import "quizz-light/pkg/quiz"

func main() {
	controller := quiz.NewController("./resources")
	controller.Run()
}
