package note

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/log"
	"yac-go/model/note"
	"yac-go/ydb"
)

func Query(d *db.DB, year, month, day, bookId int) ([]Response, error) {
	col := d.Use(ydb.ColNotes)

	start, end := calcRange(year, month, day)
	rawQuery := fmt.Sprintf(`{"int-from": %d, "int-to": %d, "in": ["key"]}`, start, end)
	if bookId != 0 {
		rawQuery = fmt.Sprintf(`{
			"n": [
				%s,
				{ "in": ["book"], "eq": %d }
			]
		}`, rawQuery, bookId)
	}
	query := []byte(rawQuery)

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

func calcRange(year, month, day int) (start, end int) {
	start = year * 10000
	if month != 0 {
		start += month * 100
	}
	if month != 0 && day != 0 {
		start += day
	}

	if month == 0 && day == 0 {
		end = year*10000 + 1332
	} else if month != 0 && day == 0 {
		end = year*10000 + month*100 + 32
	} else {
		end = start
	}

	log.Debug("query start", start)
	log.Debug("query end", end)

	return
}
