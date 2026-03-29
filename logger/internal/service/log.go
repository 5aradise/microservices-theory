package service

import (
	"context"
	"micro/logger/internal/model"
	"micro/logger/internal/storage/mongodb"
)

type Log struct {
	stor *mongodb.Entries
}

func NewLog(stor *mongodb.Entries) *Log {
	return &Log{
		stor: stor,
	}
}

func (s *Log) WriteLog(ctx context.Context, e model.Entry) error {
	_, err := s.stor.CreateLogEntry(ctx, e)
	return err
}
