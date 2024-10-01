package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/myapp/core"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "submit your response to the quizz",
	Long: `
	Submit your test by executing this command.
	First, add the player's name.
	Then, note that your answers should be formatted as '1-A',
	where '1' represents the question ID and 'A' represents your answer to the question.
	Your command should look like : go run main.go <username> 1-A 2-B 3-C`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Submit(args)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
}

func Submit(args []string) error {
	if len(args) != 4 {
		return errors.New("error : usage : go run main.go <username> 1-A 2-B 3-C")
	}

	answers := args[1:]
	userAnswers := core.UserAnswer{
		UserName: args[0],
	}

	for _, arg := range answers {
		split := strings.Split(arg, "-")

		if len(split) != 2 {
			return errors.New("error : usage : go run main.go <username> 1-A 2-B 3-C")
		}

		questionId, err := strconv.Atoi(split[0])
		if err != nil {
			return err
		}

		userAnswers.Answer = append(userAnswers.Answer, struct {
			QuestionId int
			Answer     string
		}{
			QuestionId: questionId,
			Answer:     split[1],
		})
	}

	jsonData, err := json.Marshal(userAnswers)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/submit", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(buf))
	return nil
}
