package main

import (
	"dubravaj/quiz/quiz"
	"flag"
	"time"
)

func main() {

	// create argument for input file
	csvFilename := flag.String("filename", "quiz.csv", "quiz filename")
	//create argument for quiz timeout
	quizTimeout := flag.Int("timeout", 30, "timeout for quiz to finish it")
	flag.Parse()

	// create timer for quiz
	quizTimer := time.NewTimer(time.Duration(*quizTimeout) * time.Second)

	quizQuestions := quiz.ReadQuiz(*csvFilename)
	quiz.EvaluateQuiz(quizQuestions, *quizTimer)

}
