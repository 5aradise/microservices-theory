package adapter

import (
	"micro/logger/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Entry struct {
	ID        bson.ObjectID `bson:"_id"`
	Name      string        `bson:"name"`
	Data      string        `bson:"data"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

func EntryToDB(e model.Entry) Entry {
	now := time.Now()
	return Entry{
		ID:        bson.NewObjectID(),
		Name:      e.Name,
		Data:      e.Data,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func DBToEntry(db Entry) model.Entry {
	return model.Entry{
		ID:        db.ID.Hex(),
		Name:      db.Name,
		Data:      db.Data,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
	}
}

func DBToEntries(dbs []Entry) []model.Entry {
	es := make([]model.Entry, 0, len(dbs))
	for _, db := range dbs {
		es = append(es, DBToEntry(db))
	}
	return es
}
