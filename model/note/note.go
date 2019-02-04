package note

import (
	"time"
)

type Note struct {
	Day     int      `json:"day" binding:"required"`
	Month   int      `json:"month" binding:"required"`
	Year    int      `json:"year" binding:"required"`
	Entries []*Entry `json:"entries" binding:"required"`
}

func EmptyNote() *Note {
	note := &Note{
		Day:   time.Now().Day(),
		Month: int(time.Now().Month()),
		Year:  time.Now().Year(),
	}

	return note
}
