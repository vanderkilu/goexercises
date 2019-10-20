package quiz
import (
	"bufio"
	"fmt"
	"os"
	"encoding/csv"
	"strings"
	"time"
)

type Question struct {
	question string
	answer string
}


func (q Question) ask(userChan chan string, timer <-chan time.Time) (int,int) {
	rightAnswers := 0
	totalQuestions := 0
	quest := q.question + " ="
	fmt.Println(quest)
	select {
		case <-timer: 
			return -1, -1
		case userAnswer := <-userChan:
			if strings.Compare(strings.Trim(strings.ToLower(userAnswer), "\n"), q.answer ) == 0 {
				rightAnswers += 1
			}
	}
	
	totalQuestions += 1
	return rightAnswers, totalQuestions 
}

func Main() {
	rightAnswers := 0
	totalQuestions := 0

	cwd, _ := os.Getwd()
	filePath := cwd + "/quiz/quiz.csv"
	userChan := make(chan string)
	timer := time.NewTimer(time.Duration(10) * time.Second)

	questions, err := readCsv(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	right, total := askQuestion(questions, userChan, timer.C)
	
	rightAnswers = right
	totalQuestions = total
	fmt.Printf("you had %d/%d\n right answers",rightAnswers, totalQuestions)
	
}

func readCsv(filePath string) ([]Question, error) {
	content, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	csvContent, err := csv.NewReader(content).ReadAll()

	if err != nil {
		return nil, err
	}
	questions := []Question{}

	for _, c := range csvContent {
		question := Question{c[0], c[1]}
		questions = append(questions, question)
	}
	return questions, nil

}

func askQuestion(questions []Question, 
	userChan chan string, 
	timer <-chan time.Time) (int, int) {
	r := 0
	t := 0

	go getUserInput(userChan)

	for _, quest := range questions {
		right, total := quest.ask(userChan, timer)
		if (right == -1) {
			close(userChan)
			fmt.Println("time up")
			return r,t
		}
		r += right
		t += total
	}
	close(userChan)
	return r, t

}

func getUserInput(userChan chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		results, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		userChan <- results
	}
}