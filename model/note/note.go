package note

import (
	"time"
)

type Note struct {
	Key       int        `json:"key"`
	CreatedAt *time.Time `json:"createdAt" binding:"required"`
	Entries   []*Entry   `json:"entries" binding:"required"`
}

func EmptyNote() *Note {
	t := time.Now()

	note := &Note{
		Key:       MakeKey(&t),
		CreatedAt: &t,
	}

	return note
}

func MakeKey(t *time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}
