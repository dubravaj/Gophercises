package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	answer   string
}

func createQuestionsList(data [][]string) []Question {
	var questions []Question
	for _, questionData := range data {
		question := Question{question: questionData[0], answer: strings.TrimSpace(questionData[1])}
		questions = append(questions, question)
	}

	return questions
}

func ReadQuiz(filename string) []Question {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	// handle closing file
	defer file.Close()

	// read data from CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error while reading quiz file", err)
	}
	questionsList := createQuestionsList(records)
	return questionsList
}

func EvaluateQuiz(quizQuestions []Question, timer time.Timer) {

	scanner := bufio.NewScanner(os.Stdin)
	correctlyAnswered := 0
	questionsTotal := len(quizQuestions)
	answerChan := make(chan string)

	for _, question := range quizQuestions {
		fmt.Printf("Anwer for %s is: ", question.question)

		// run goroutine for getting answer to not block program when
		// signal from timer is received
		go func() {
			scanner.Scan()
			err := scanner.Err()
			if err != nil {
				log.Fatal(err)
			}
			answer := scanner.Text()
			answer = strings.Trim(answer, " \t")
			answerChan <- answer

		}()

		select {
		case <-timer.C:
			return
		case answer := <-answerChan:
			if question.answer == answer {
				correctlyAnswered++
			}
		}
	}

	fmt.Println()
	fmt.Println("------Quiz results------")
	fmt.Printf("Total number of questions: %d\n", questionsTotal)
	fmt.Printf("Correctly answered questions: %d\n", correctlyAnswered)
	fmt.Printf("Incorrectly answered questions: %d\n", questionsTotal-correctlyAnswered)

}
