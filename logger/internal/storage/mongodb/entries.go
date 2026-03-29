package mongodb

import (
	"context"
	"errors"
	"micro/logger/internal/model"
	"micro/logger/internal/storage/mongodb/adapter"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var createdAtDesc = bson.D{{Key: "created_at", Value: -1}}

func byID(id bson.ObjectID) bson.D {
	return bson.D{{Key: "_id", Value: id}}
}

func updateAllEntry(e model.Entry) bson.D {
	return bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "name", Value: e.Name},
			{Key: "data", Value: e.Data},
			{Key: "updated_at", Value: time.Now()},
		}},
	}
}

type Entries struct {
	logs *mongo.Collection
}

func NewEntries(cl *mongo.Client) *Entries {
	return new(Entries{
		logs: cl.Database("logs").Collection("logs"),
	})
}

func (s *Entries) CreateLogEntry(ctx context.Context, e model.Entry) (model.Entry, error) {
	db := adapter.EntryToDB(e)
	res, err := s.logs.InsertOne(ctx, db)
	if err != nil {
		return model.Entry{}, err
	}
	db.ID = res.InsertedID.(bson.ObjectID)

	return adapter.DBToEntry(db), nil
}

func (s *Entries) AllLogEntries(ctx context.Context) ([]model.Entry, error) {
	cur, err := s.logs.Find(ctx, bson.D{}, options.Find().
		SetSort(createdAtDesc))
	if err != nil {
		return nil, err
	}

	var logs []adapter.Entry
	err = cur.All(ctx, &logs)
	if err != nil {
		return nil, err
	}

	return adapter.DBToEntries(logs), nil
}

func (s *Entries) GetLogEntryByID(ctx context.Context, id string) (model.Entry, error) {
	oid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.Entry{}, err
	}

	var res adapter.Entry
	err = s.logs.FindOne(ctx, byID(oid)).Decode(&res)
	if err != nil {
		return model.Entry{}, err
	}

	return adapter.DBToEntry(res), nil
}

func (s *Entries) UpdateLogEntry(ctx context.Context, e model.Entry) error {
	id, err := bson.ObjectIDFromHex(e.ID)
	if err != nil {
		return err
	}

	res, err := s.logs.UpdateByID(ctx, id, updateAllEntry(e))
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("entry not found")
	}

	return nil
}

func (s *Entries) Drop(ctx context.Context) error {
	return s.logs.Drop(ctx)
}
