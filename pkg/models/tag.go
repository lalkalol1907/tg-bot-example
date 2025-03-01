package models

type Tag struct {
	Id     string `db:"id"`
	Text   string `db:"tag"`
	GoodId string `db:"good_id"`
}
