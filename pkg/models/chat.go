package models

type Chat struct {
	Id       int64 `db:"id"`
	WorkerId int64 `db:"worker_id"`
	OwnerId  int64 `db:"owner_id"`
}
