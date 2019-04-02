package book

import (
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/ydb"
)

func Delete(d *db.DB, id int) error {
	col := d.Use(ydb.ColBooks)
	if err := col.Delete(id); err != nil {
		return err
	}

	return nil
}
