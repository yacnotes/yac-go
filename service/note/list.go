package note

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/model/note"
	"yac-go/ydb"
)

func GetAll(d *db.DB, bookId int) ([]Response, error) {
	col := d.Use(ydb.ColNotes)

	query := []byte(`["all"]`)
	if bookId != 0 {
		query = []byte(fmt.Sprintf(`{ "in": ["book"], "eq": %d }`, bookId))
	}
	queryResult, err := ydb.ExecuteQuery(col, query)
	if err != nil {
		return nil, err
	}

	i := 0
	notes := make([]Response, len(queryResult))
	for id := range queryResult {
		n := &note.Note{}
		if err := ydb.LoadById(col, id, n); err != nil {
			return nil, err
		}
		notes[i] = Response{
			Id:   id,
			Note: n,
		}
		i++
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
