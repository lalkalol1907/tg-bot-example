package models

type Worker struct {
	Id      int64 `db:"id"`
	OwnerId int64 `db:"owner_id"`
}
