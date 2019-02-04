package note

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/log"
	"yac-go/model"
	"yac-go/model/note"
)

func GetAll(d *db.DB, page, total int) map[int]*note.Note {
	col := d.Use("notes")

	queryResult := make(map[int]struct{})
	var query interface{}
	if err := json.Unmarshal([]byte(`["all"]`), &query); err != nil {
		log.Panic("Failed to build query:", err)
	}

	if err := db.EvalQuery(query, col, &queryResult); err != nil {
		log.Panic("Failed to run query:", err)
	}

	notes := make(map[int]*note.Note)

	for id := range queryResult {
		plain, err := col.Read(id)
		if err != nil {
			log.Panic("Failed to read from query result")
		}

		n := &note.Note{}
		model.Marshal(plain, n)
		notes[id] = n
	}

	return notes
}
