package book

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/model"
	"yac-go/model/book"
	"yac-go/ydb"
)

func Add(d *db.DB, b *book.Book) (int, error) {
	col := d.Use(ydb.ColBooks)

	id, err := col.Insert(model.Unmarshal(b))
	if err != nil {
		return 0, err
	}

	return id, nil
}
