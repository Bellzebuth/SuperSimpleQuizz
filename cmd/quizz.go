package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/myapp/core"
)

var quizzCmd = &cobra.Command{
	Use:   "quizz",
	Short: "request API  to get a quizz",
	Long:  "usage : go run main.go quizz",
	Run: func(cmd *cobra.Command, args []string) {
		err := quizz()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(quizzCmd)
}

func quizz() error {
	req, err := http.NewRequest("GET", "http://localhost:8080/quizz", nil)
	if err != nil {
		return err
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var quizz []core.QuestionWithoutAnswer

	if resp.Body != nil {
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(buf, &quizz)
		if err != nil {
			return err
		}
	}

	for _, question := range quizz {
		fmt.Printf("%d - %s\n", question.Id, question.Question)
		for _, option := range question.Options {
			fmt.Println(option)
		}
		fmt.Println()
	}

	return nil
}
