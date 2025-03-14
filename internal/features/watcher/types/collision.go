package types

import tag_check "bot-test/pkg/tag-check"

type Collision struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	OwnerId int64 `json:"owner_id"`

	Result []*tag_check.FindTagsResult `json:"result"`
}
