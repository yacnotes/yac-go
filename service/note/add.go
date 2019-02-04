package note

import (
	"errors"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/log"
	"yac-go/model"
	"yac-go/model/note"
	"yac-go/ydb"
)

func Add(d *db.DB, n *note.Note) (int, int, error) {
	col := d.Use(ydb.ColNotes)

	query := fmt.Sprintf(`{
		"n": [
			{ "in": ["day"], "eq": %d },
			{ "in": ["month"], "eq": %d },
			{ "in": ["year"], "eq": %d }
		]
	}`, n.Day, n.Month, n.Year)

	queryResult, err := ydb.ExecuteQuery(col, []byte(query))
	if err != nil {
		log.Panic("Failed to execute query:", query, err)
	}

	if len(queryResult) != 0 {
		return 0, getId(queryResult), errors.New("already have note for that day")
	}

	injectEntryIds(n)

	id, err := col.Insert(model.Unmarshal(n))
	if err != nil {
		return 0, 0, err
	}

	return id, 0, err
}

func getId(m map[int]struct{}) int {
	for id := range m {
		return id
	}
	return 0
}

func injectEntryIds(n *note.Note) {
	for _, entry := range n.Entries {
		entry.ID = model.GenerateId()
	}
}
