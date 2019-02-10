package book

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/model/book"
	"yac-go/ydb"
)

func Get(d *db.DB, id int) (*book.Book, error) {
	col := d.Use(ydb.ColBooks)

	b := &book.Book{}
	if err := ydb.LoadById(col, id, b); err != nil {
		return nil, err
	}

	return b, nil
}