package note

import (
	"time"
)

type Note struct {
	Day     time.Time `json:"day"`
	Entries []*Entry  `json:"entries"`
}

func EmptyNote() *Note {
	note := &Note{
		Day: time.Now(),
	}

	return note
}
