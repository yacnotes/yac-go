package note

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/log"
	"yac-go/model/note"
	"yac-go/ydb"
)

// GetById tries to fetch the note object from a database id
func GetById(d *db.DB, id int) (*note.Note, error) {
	col := d.Use(ydb.ColNotes)

	n := &note.Note{}
	if err := ydb.LoadById(col, id, n); err != nil {
		return nil, err
	}

	return n, nil
}

// GetByKey tries to find one note with a given key (and optionally book, if bookId != 0)
func GetByKey(d *db.DB, key int, bookId string) (*note.Note, int) {
	col := d.Use(ydb.ColNotes)

	query := fmt.Sprintf(`{ "in": ["key"], "eq": %d, "limit": 1 }`, key)
	if bookId != "" {
		query = fmt.Sprintf(`{
			"n": [
				{ "in": ["key"], "eq": %d },
				{ "in": ["book"], "eq": "%s" }
			],
			"limit": 1
		}`, key, bookId)
	}

	queryResult, err := ydb.ExecuteQuery(col, []byte(query))
	if err != nil {
		log.Panic("Failed to execute query:", query, err)
	}

	if len(queryResult) != 1 {
		return nil, 0
	}

	id := getId(queryResult)

	n := &note.Note{}
	if err := ydb.LoadById(col, id, n); err != nil {
		log.Panic("Failed to load note by id (GetByKey)")
	}

	return n, id
}

func getId(m map[int]struct{}) int {
	for id := range m {
		return id
	}
	return 0
}
