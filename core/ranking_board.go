package core

import (
	"fmt"
	"net/http"
)

type Score struct {
	Player string
	Score  int
}

type RankingBoard struct {
	scores []Score
}

func (rb *RankingBoard) InsertScore(newScore Score) int {
	position := 0
	for i, s := range rb.scores {
		if newScore.Score > s.Score {
			position = i
			break
		} else {
			position = i + 1
		}
	}

	rb.scores = append(rb.scores[:position], append([]Score{newScore}, rb.scores[position:]...)...)

	return position + 1
}

func (rb *RankingBoard) ToString() string {
	var rankingBoard string
	for i, score := range rb.scores {
		rankingBoard = fmt.Sprintf("%s%d - player : %s, score : %d\n",
			rankingBoard,
			i+1,
			score.Player,
			score.Score)
	}

	return rankingBoard
}

func ShowRankingBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rankingBoard.ToString())
}
