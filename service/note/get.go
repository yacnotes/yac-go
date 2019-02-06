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

// GetById tries to fetch the note object from a database id
func GetById(d *db.DB, id int) (*note.Note, error) {
	col := d.Use(ydb.ColNotes)

	res, err := col.Read(id)
	if err != nil {
		return nil, err
	}

	n := &note.Note{}
	model.Marshal(res, &n)

	return n, nil
}

// GetByKey tries to find the note with a given key
func GetByKey(d *db.DB, key int) (*note.Note, int) {
	col := d.Use(ydb.ColNotes)

	query := fmt.Sprintf(`{ "in": ["key"], "eq": %d, "limit": 1 }`, key)

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

// GetByIdentifier tries to find a note by its key or id; returns note and id
func GetByIdentifier(d *db.DB, identifier int) (*note.Note, error, int) {
	var n *note.Note
	var id int
	if identifier < 100000000 {
		// got key
		n, id = GetByKey(d, identifier)
	} else {
		// got id
		n, _ = GetById(d, identifier)
		id = identifier
	}

	if n == nil {
		return nil, errors.New("note not found"), 0
	}

	return n, nil, id
}

func getId(m map[int]struct{}) int {
	for id := range m {
		return id
	}
	return 0
}
