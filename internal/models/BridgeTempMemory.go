package models

import "sync"

type BridgeTempMemory struct {
	Timestamp int64
	RelayName string
	MessageDs []MessageDs
	MessageTg []MessageTg
	Wg        sync.WaitGroup
}
type MessageDs struct {
	MessageId string
	ChatId    string
}
type MessageTg struct {
	MessageId string
	ChatId    string
}
