package core

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Question struct {
	Id       int
	Question string
	Options  []string
	Answer   string
}

var questions = []Question{
	{
		Id:       1,
		Question: "What is the capital of France?",
		Options:  []string{"A. Berlin", "B. Madrid", "C. Paris", "D. Lisbon"},
		Answer:   "C",
	},
	{
		Id:       2,
		Question: "Which planet is known as the Red Planet?",
		Options:  []string{"A. Earth", "B. Mars", "C. Jupiter", "D. Venus"},
		Answer:   "B",
	},
	{
		Id:       3,
		Question: "What is the chemical symbol for water?",
		Options:  []string{"A. H2O", "B. O2", "C. CO2", "D. HO"},
		Answer:   "A",
	},
	{
		Id:       4,
		Question: "How many continents are there on Earth?",
		Options:  []string{"A. 5", "B. 6", "C. 7", "D. 8"},
		Answer:   "C",
	},
	{
		Id:       5,
		Question: "Which is the largest ocean on Earth?",
		Options:  []string{"A. Atlantic Ocean", "B. Indian Ocean", "C. Pacific Ocean", "D. Arctic Ocean"},
		Answer:   "C",
	},
	{
		Id:       6,
		Question: "Who wrote 'Romeo and Juliet'?",
		Options:  []string{"A. Charles Dickens", "B. J.K. Rowling", "C. William Shakespeare", "D. Mark Twain"},
		Answer:   "C",
	},
	{
		Id:       7,
		Question: "What is the smallest prime number?",
		Options:  []string{"A. 0", "B. 1", "C. 2", "D. 3"},
		Answer:   "C",
	},
	{
		Id:       8,
		Question: "Which country is known for the maple leaf symbol?",
		Options:  []string{"A. Australia", "B. Germany", "C. Canada", "D. New Zealand"},
		Answer:   "C",
	},
	{
		Id:       9,
		Question: "In computing, what does 'CPU' stand for?",
		Options:  []string{"A. Central Processing Unit", "B. Central Power Unit", "C. Computer Personal Unit", "D. Central Program Unit"},
		Answer:   "A",
	},
	{
		Id:       10,
		Question: "Which element is represented by the symbol 'Fe'?",
		Options:  []string{"A. Lead", "B. Iron", "C. Gold", "D. Silver"},
		Answer:   "B",
	},
}

// To avoid bothering the user, we will send different quizzes.
// To do this, we copy the questions, shuffle them, and then
// select the first n questions.
func getRandomQuestions(n int) []Question {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	questionsSlice := make([]Question, 0, len(questions))
	questionsSlice = append(questionsSlice, questions...)

	r.Shuffle(len(questionsSlice), func(i, j int) {
		questionsSlice[i], questionsSlice[j] = questionsSlice[j], questionsSlice[i]
	})

	if n > len(questionsSlice) {
		n = len(questionsSlice)
	}

	return questionsSlice[:n]
}

// We need this struct to not send the answer to players
type QuestionWithoutAnswer struct {
	Id       int
	Question string
	Options  []string
}

func removeAnswers(questions []Question) []QuestionWithoutAnswer {
	var result []QuestionWithoutAnswer
	for _, q := range questions {
		result = append(result, QuestionWithoutAnswer{
			Id:       q.Id,
			Question: q.Question,
			Options:  q.Options,
		})
	}
	return result
}

func QuizzHandler(w http.ResponseWriter, r *http.Request) {
	// get 3 random questions
	randomQuestions := getRandomQuestions(3)

	// remove answer from the questions
	questionsWithoutAnswers := removeAnswers(randomQuestions)

	// send quizz
	json.NewEncoder(w).Encode(questionsWithoutAnswers)
}

/// questions

var rankingBoard = &RankingBoard{}

type Answer struct {
	QuestionId int
	Answer     string
}

type UserAnswer struct {
	UserName string
	Answer   []Answer
}

func CheckAnswersHandler(w http.ResponseWriter, r *http.Request) {
	// get request body
	var userAnswer UserAnswer
	err := json.NewDecoder(r.Body).Decode(&userAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// verify answers
	score := Score{
		Player: userAnswer.UserName,
		Score:  0,
	}
	for _, answer := range userAnswer.Answer {
		for _, question := range questions {
			if question.Id == answer.QuestionId {
				if question.Answer == answer.Answer {
					score.Score += 3
				} else {
					score.Score -= 1
				}
			}
		}
	}

	// insert player in the ranking board
	position := rankingBoard.InsertScore(score)

	// set message
	var message string
	if len(rankingBoard.scores) == 1 {
		message = "You are the first player !"
	} else if position == 1 {
		message = "You are first !"
	} else if position == len(rankingBoard.scores) {
		message = "You are the last..."
	} else {
		percentage := -(float64(position-len(rankingBoard.scores)) / float64(len(rankingBoard.scores))) * 100

		message = fmt.Sprintf("You are better than %.0f%% of all quizzers", percentage)
	}

	// send response
	fmt.Fprintf(w, "Your score is %d\nYou are ranked %d out of %d\n%s",
		score.Score,
		position,
		len(rankingBoard.scores),
		message)
}
