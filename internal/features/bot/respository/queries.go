package respository

import _ "embed"

var (
	//go:embed queries/add_good.sql
	AddGoodQuery string

	//go:embed queries/delete_good.sql
	DeleteGoodQuery string

	//go:embed queries/get_goods_by_owner_id.sql
	GetGoodsByOwnerIdQuery string
)
