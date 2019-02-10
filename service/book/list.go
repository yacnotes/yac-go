package book

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/model/book"
	"yac-go/ydb"
)

func GetAll(d *db.DB) ([]Response, error) {
	col := d.Use(ydb.ColBooks)

	query := []byte(`["all"]`)
	queryResult, err := ydb.ExecuteQuery(col, query)
	if err != nil {
		return nil, err
	}

	i := 0
	books := make([]Response, len(queryResult))
	for id := range queryResult {
		b := &book.Book{}
		if err := ydb.LoadById(col, id, b); err != nil {
			return nil, err
		}
		response := &Response{
			Id:   id,
			Book: b,
		}
		books[i] = *response
		i++
	}

	return books, nil
}
