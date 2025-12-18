package kanban

import (
	"encoding/json"
	"os"
)

const dbPath = "db.json"

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Board [3][]Task

func (b Board) Save() {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return
	}
	_ = os.WriteFile(dbPath, data, 0644)
}

func Load() (Board, bool) {
	var board Board
	data, err := os.ReadFile(dbPath)
	if err != nil {
		return Board{
			{{Title: "Buy milk", Description: "Strawberry milk"}},
			{{Title: "Eat sushi", Description: "Rolls and sashimi"}},
			{{Title: "Fold laundry", Description: "Only the socks"}},
		}, false
	}

	err = json.Unmarshal(data, &board)
	if err != nil {
		return Board{}, false
	}
	return board, true
}
