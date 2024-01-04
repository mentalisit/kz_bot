package models

type BridgeTempMemory struct {
	Timestamp int64
	RelayName string
	MessageDs []MessageDs
	MessageTg []MessageTg
}
type MessageDs struct {
	MessageId string
	ChatId    string
}
type MessageTg struct {
	MessageId int
	ChatId    string
}
