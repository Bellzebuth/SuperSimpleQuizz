package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomQuestions(t *testing.T) {
	assert := assert.New(t)

	randomQuestions := getRandomQuestions(3)
	assert.Len(randomQuestions, 3)

	randomQuestions = getRandomQuestions(1)
	assert.Len(randomQuestions, 1)

	randomQuestions = getRandomQuestions(10)
	assert.Len(randomQuestions, 10)

	randomQuestions = getRandomQuestions(100)
	assert.Len(randomQuestions, 10)
}

func TestCheckAnswersHandler(t *testing.T) {
	userAnswer := UserAnswer{
		UserName: "John",
		Answer: []Answer{
			{QuestionId: 1, Answer: "C"},
			{QuestionId: 2, Answer: "C"},
			{QuestionId: 3, Answer: "A"},
		},
	}

	body, err := json.Marshal(userAnswer)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/check-answers", bytes.NewBuffer(body))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(CheckAnswersHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.Equal(t, "Your score is 5\nYou are ranked 1 out of 1\nYou are the first player !", recorder.Body.String())
}
