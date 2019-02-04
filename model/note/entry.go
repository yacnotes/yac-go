package note

import (
	"time"
	"yac-go/model"
)

type Entry struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ChangedAt time.Time `json:"changedAt"`
	Text      string    `json:"text"`
	Tags      []string  `json:"tags"`
}

func EmptyEntry() *Entry {
	now := time.Now()

	entry := &Entry{
		ID:        model.GenerateId(),
		CreatedAt: now,
		ChangedAt: now,
	}

	return entry
}
