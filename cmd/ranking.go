package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// rankingCmd represents the ranking command
var rankingCmd = &cobra.Command{
	Use:   "ranking",
	Short: "print the ranking board",
	Long:  "usage : go run main.go ranking",
	Run: func(cmd *cobra.Command, args []string) {
		Ranking()
	},
}

func init() {
	rootCmd.AddCommand(rankingCmd)
}

func Ranking() error {
	req, err := http.NewRequest("GET", "http://localhost:8080/ranking", nil)
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
