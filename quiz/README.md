## Simple quiz 

This program is intended to read quiz defined in CSV file in format "question, answer".

Quiz questions are asked by the program and user should provide answer.

Program also support setup of timeout for quiz to be finished, e.g. if timeout is 30 seconds, user have to answer questions in this time limit, otherwise after timeout, quiz is immediately stopped and evaluated.

Usage:

go run main.go -filename problems.csv -timeout 30
