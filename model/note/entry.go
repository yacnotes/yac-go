package note

import (
	"time"
	"yac-go/model"
)

type Entry struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	ChangedAt time.Time `json:"changedAt" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	Tags      []string  `json:"tags" binding:"required"`
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
