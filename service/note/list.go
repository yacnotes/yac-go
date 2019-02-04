package note

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/model/note"
	"yac-go/ydb"
)

func GetAll(d *db.DB) (map[int]*note.Note, error) {
	col := d.Use(ydb.ColNotes)

	queryResult, err := ydb.ExecuteQuery(col, []byte(`["all"]`))
	if err != nil {
		return nil, err
	}

	notes := make(map[int]*note.Note)
	for id := range queryResult {
		n := &note.Note{}
		if err := ydb.LoadById(col, id, n); err != nil {
			return nil, err
		}
		notes[id] = n
	}

	return notes, nil
}

func GetAllIds(d *db.DB) ([]int, error) {
	col := d.Use(ydb.ColNotes)

	res := make(map[int]struct{})
	if err := db.EvalAllIDs(col, &res); err != nil {
		return nil, err
	}

	i := 0
	allIds := make([]int, len(res))
	for id := range res {
		allIds[i] = id
		i++
	}

	return allIds, nil
}
