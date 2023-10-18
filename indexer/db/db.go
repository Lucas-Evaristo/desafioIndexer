package db

import "indexador/model"

type Database interface {
	InsertRecords(model.BulkEmail) error
}
