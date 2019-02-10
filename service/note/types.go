package note

import "yac-go/model/note"

type Response struct {
	Id   int        `json:"id"`
	Note *note.Note `json:"note"`
}
