package models

type Good struct {
	Id      string `db:"id"`
	Name    string `db:"name"`
	OwnerId int64  `db:"owner_id"`
}
