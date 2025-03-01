package types

type IncomingMessage struct {
	WorkerId  int64
	ChatId    int64
	UserId    int64
	MessageId int64
	Text      string
}
