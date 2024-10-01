package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/myapp/core"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the HTTP server",
	Long:  "usage : go run main.go server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func startServer() {
	http.HandleFunc("/ranking", core.ShowRankingBoard)
	http.Handle("/quizz", http.HandlerFunc(core.QuizzHandler))
	http.Handle("/submit", http.HandlerFunc(core.CheckAnswersHandler))

	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
