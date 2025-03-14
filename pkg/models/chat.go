package models

type Chat struct {
	Id      int64 `db:"id"`
	OwnerId int64 `db:"owner_id"`
}
