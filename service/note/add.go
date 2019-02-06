package note

import (
	"errors"
	"github.com/HouzuoGuo/tiedot/db"
	"yac-go/model"
	"yac-go/model/note"
	"yac-go/ydb"
)

func Add(d *db.DB, n *note.Note) (int, int, error) {
	col := d.Use(ydb.ColNotes)

	n.Key = note.MakeKey(n.CreatedAt)

	oldNote, oldId := GetByKey(d, n.Key)
	if oldNote != nil {
		return 0, oldId, errors.New("already have note for that day")
	}

	injectEntryIds(n)

	id, err := col.Insert(model.Unmarshal(n))
	if err != nil {
		return 0, 0, err
	}

	return id, 0, err
}

func injectEntryIds(n *note.Note) {
	for _, entry := range n.Entries {
		entry.ID = model.GenerateId()
	}
}
