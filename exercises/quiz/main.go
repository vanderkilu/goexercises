package main	


import (
	"os"
	"encoding/csv"
	"log"
	"io"
	"bufio"
	"strings"
	"fmt"
)

type Question struct {
	question string
	answer string
}

func readQuestions() []Question {

	file, err := os.Open("./quiz.csv")
	if err != nil {
		log.Fatalln("error opening the quiz file")
	}

	csvFile := csv.NewReader(file)
	var questions []Question

	for {
		csvRecord, err := csvFile.Read()
		if err == io.EOF {
			break
		}
		questions = append(questions, Question {
			question: csvRecord[0],
			answer: csvRecord[1], 
		})
	}
	return questions
}

func getUserInputs(userInputs chan string) {
	
	reader := bufio.NewReader(os.Stdin)

	for {
		userAnswer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("can't read response from terminal")
		}
		userAnswer = strings.Replace(userAnswer, "\n", "", -1)
		userInputs <- userAnswer
	}
}

func checkAnswer(question Question, results chan string) int {
	
	fmt.Printf("%s = \n", question.question)
	userAnswer := <- results
	if strings.Compare(userAnswer, question.answer) == 0 {
		return 1
	}
	return 0
}

func main() {
	results := make(chan string)
	questions := readQuestions()

	go getUserInputs(results)

	correctAnswers := 0
	totalQuestions := 0

	fmt.Println("welcome to the terminal quiz exercise")

	for _, question := range questions {
		correctAnswers += checkAnswer(question, results)
		totalQuestions += 1
	}
	
	
}