package models

type BridgeTempMemory struct {
	Timestamp int64
	RelayName string
	MessageDs []struct {
		MessageId string
		ChatId    string
	}
	MessageTg []struct {
		MessageId int
		ChatId    int64
	}
}
